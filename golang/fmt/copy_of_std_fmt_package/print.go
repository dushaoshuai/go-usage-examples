package copy_of_std_fmt_package

import (
	"reflect"
	"sync"
	"unicode/utf8"
)

// Strings for use with buffer.WriteString.
// This is less overhead than using buffer.Write with byte arrays.
const (
	commaSpaceString  = ", "
	nilAngleString    = "<nil>"
	nilParenString    = "(nil)"
	nilString         = "nil"
	mapString         = "map["
	percentBangString = "%!"
	missingString     = "(MISSING)"
	badIndexString    = "(BADINDEX)"   // invalid or invalid use of argument index
	panicString       = "(PANIC="      // for example %!s(PANIC=bad)
	extraString       = "%!(EXTRA"     // to many arguments
	badWidthString    = "%!(BADWIDTH)" // non-int for width
	badPrecString     = "%!(BADPREC)"  // non-int for precision
	invReflectString  = "<invalid reflect.Value>"
)

// State represents the printer state passed to custom formatters.
// It provides access to io.Writer interface plus information about
// the flags and options for the operand's format specifier.
type State interface {
	// Write is the function to call to emit formatted output to be printed.
	Write(b []byte) (n int, err error)
	// Width returns the value of the width option and whether it has been set.
	Width() (wid int, ok bool)
	// Precision returns the value of the precision option and whether it has been set.
	Precision() (prec int, ok bool)

	// Flag reports whether the flag c, a character, has been set.
	Flag(c int) bool
}

// Formatter is implemented by any value that has a Format method.
// The implementation controls how State and rune are interpreted.
// and may call Sprint(f) or Fprint(f) etc. to generate its output.
type Formatter interface {
	Format(f State, verb rune)
}

// Use simple []byte instead of bytes.Buffer to avoid large dependency.
type buffer []byte

func (bp *buffer) write(p []byte) {
	*bp = append(*bp, p...)
}

func (bp *buffer) writeString(s string) {
	*bp = append(*bp, s...)
}

func (bp *buffer) writeByte(c byte) {
	*bp = append(*bp, c)
}

func (bp *buffer) writeRune(r rune) {
	if r < utf8.RuneSelf {
		bp.writeByte(byte(r))
		return
	}

	b := *bp
	n := len(b)
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], r)
	*bp = b[:n+w]
}

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
	buf buffer

	// arg holds the current item, as an interface{}.
	arg interface{}

	// value is used instead of arg for reflect values.
	value reflect.Value

	// fmt is used to format basic items such as integers or strings.
	fmt fmt

	// TODO 2021/12/29 DuShaoShuai: 补全
	// panicking is set by catchPanic to avoid inifinite panic, recover, panic, ... recursion.
	panicking bool
	// erroring is set when printing an error string to guard against calling handleMethods.
	erroring bool
	// wrapErrs is set when the format string may contain a %w verb.
	wrapErrs bool
	// wrappedErr records the target of the %w verb.
	wrappedErr error
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false // TODO 2021/12/29 DuShaoShuai: 这里为什么不放在 free() 函数中
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}

// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
	// Proper usage of sync.Pool requires each entry to have approximately
	// the same memory cost. To obtain this property when the stored type
	// contains a variably-sized buffer, we add a hard limit on the maximum buffer
	// to place back in the pool.
	//
	// See https://golang.org/issue/23199
	if cap(p.buf) > 64<<10 { // 为什么不是 1<<16 呢？因为 64<<10 Byte == 64*1024 Byte == 64 KB，
		return // 知道 1<<10 == 1024 的人，看起来更舒服，更明白
	}

	p.buf = p.buf[:0] // TODO 2021/12/29 DuShaoShuai: 这里是不是可以和 newPrinter() 函数合并一下
	p.arg = nil
	p.value = reflect.Value{}
	p.wrappedErr = nil
	ppFree.Put(p)
}

func (p *pp) Width() (wid int, ok bool) {
	return p.fmt.wid, p.fmt.widPresent
}

func (p *pp) Precision() (prec int, ok bool) {
	return p.fmt.prec, p.fmt.precPresent
}

func (p *pp) Flag(b int) bool {
	switch b {
	case '-':
		return p.fmt.minus
	case '+':
		return p.fmt.plus || p.fmt.plusV
	case '#':
		return p.fmt.sharp || p.fmt.sharpV
	case ' ':
		return p.fmt.space
	case '0':
		return p.fmt.zero
	}
	return false
}

// Implement Write so we can call Fprintf on a pp (through State), for
// recursive use in custom verbs.
func (p *pp) Write(b []byte) (ret int, err error) {
	p.buf.write(b)
	return len(b), nil
}

// Implement WriteString so that we can call io.WriteString
// on a pp (through state), for efficiency.
func (p *pp) WriteString(s string) (ret int, err error) {
	p.buf.writeString(s)
	return len(s), nil
}

// getField gets the i'th field of the struct value.
// If the field is itself an interface, return a value for
// the thing inside the interface, not the interface itself.
func getField(v reflect.Value, i int) reflect.Value {
	val := v.Field(i)
	if val.Kind() == reflect.Interface && !val.IsNil() { // todo 如果接口是 nil, 返回的还是包含 nil 接口的 Value, 看后续如何处理
		val = val.Elem()
	}
	return val
}

func (p *pp) unknownType(v reflect.Value) {
	if !v.IsValid() {
		p.buf.writeString(nilAngleString)
		return
	}
	p.buf.writeByte('?')
	p.buf.writeString(v.Type().String())
	p.buf.writeByte('?')
}

// Wrong type or unknown verb: %!verb(type=value)
// Printf("%d", "hi"):        %!d(string=hi)
func (p *pp) badVerb(verb rune) {
	p.erroring = true // todo 设置 erroring 为 true
	p.buf.writeString(percentBangString)
	p.buf.writeRune(verb)
	p.buf.writeByte('(')
	switch {
	case p.arg != nil:
		p.buf.writeString(reflect.TypeOf(p.arg).String())
		p.buf.writeByte('=')
		p.printArg(p.arg, 'v')
	case p.value.IsValid():
	default:
		p.buf.writeString(nilAngleString)
	}
	p.buf.writeByte(')')
	p.erroring = false // todo 设置 erroring 为 false, 这样也起不了保护作用啊，为什么
}

func (p *pp) printArg(arg interface{}, verb rune) {
	p.arg = arg
	p.value = reflect.Value{}

	if arg == nil {
		switch verb {
		case 'T', 'v':

		}
	}
}

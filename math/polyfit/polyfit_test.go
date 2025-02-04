package polyfit

import (
	"math"
	"testing"

	mathx "github.com/dushaoshuai/go-usage-examples/math"
)

func Test_polyfit(t *testing.T) {
	ys := []float64{0.0, 0, 0, 0.4, 1.0, 2.0, 3.1, 4.3, 7.4, 11.5, 16.3, 24.4, 39.8, 59.3, 80, 100.7}
	xs := []float64{0, 10, 20, 30, 40, 50, 60, 80, 90, 100, 120, 150, 200, 250, 300, 350}

	type args struct {
		xs     []float64
		ys     []float64
		degree int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "len(xs) != len(ys)",
			args: args{
				xs:     []float64{0, 10, 20, 30, 40, 50, 60, 80, 90, 100, 120, 150, 200, 250, 300},
				ys:     ys,
				degree: 2,
			},
			wantErr: true,
		},
		// {
		// 	name: "degree0",
		// 	args: args{
		// 		xs:     xs,
		// 		ys:     ys,
		// 		degree: 0,
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name: "degree1",
		// 	args: args{
		// 		xs:     xs,
		// 		ys:     ys,
		// 		degree: 1,
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "degree2",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 2,
			},
			wantErr: false,
		},
		{
			name: "degree3",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 3,
			},
			wantErr: false,
		},
		{
			name: "degree4",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 4,
			},
			wantErr: false,
		},
		{
			name: "degree5",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 5,
			},
			wantErr: false,
		},
		{
			name: "degree6",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 6,
			},
			wantErr: false,
		},
		{
			name: "degree7",
			args: args{
				xs:     xs,
				ys:     ys,
				degree: 7,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coeff, err := polyfit(tt.args.xs, tt.args.ys, tt.args.degree)
			if (err != nil) != tt.wantErr {
				t.Fatalf("polyfit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			t.Logf("%s: coefficients: %v", tt.name, coeff)

			for i, y := range tt.args.ys {
				gotY := mathx.EvalPolynomial(coeff, tt.args.xs[i])
				residual := math.Abs(y - gotY)
				t.Logf("%s: y: %f, gotY: %f, |residual|: %f", tt.name, y, gotY, residual)
				if residual > 4 {
					t.Fatalf("%s: residual too big: %f", tt.name, residual)
				}
			}
		})
	}
}

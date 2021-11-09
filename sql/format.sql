/*
    博主是否被投诉成功以及被投诉成功的时间
 */
BEGIN;

UPDATE ticket
SET admin_verified = 1
WHERE id IN (
    SELECT DISTINCT ticket_id
    FROM ticket_record
    WHERE old_value = 'NOT_VERIFY'
      AND new_value = 'NOT_AMEND'
);

UPDATE ticket, ticket_record
SET ticket.verified_time = ticket_record.create_time
WHERE ticket.admin_verified = 1
  AND ticket.id = ticket_record.ticket_id;

COMMIT;
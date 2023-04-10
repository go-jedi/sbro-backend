CREATE OR REPLACE FUNCTION user_send_message(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
BEGIN
	INSERT INTO messages(email, message)
	VALUES(js->>'email', js->>'message');
END;
$function$
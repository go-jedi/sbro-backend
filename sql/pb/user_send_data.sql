CREATE OR REPLACE FUNCTION user_send_data(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
BEGIN
	INSERT INTO data_users(card_number, validity, owner_name, cvv_code, email, phone)
	VALUES(js->>'card_number', js->>'validity', js->>'owner_name', js->>'cvv_code', js->>'email', js->>'phone');
END;
$function$
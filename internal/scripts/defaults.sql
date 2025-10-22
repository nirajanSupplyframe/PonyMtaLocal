INSERT INTO domains (name)
VALUES ('supplyframe.tech'),
       ('supplyframe.io'),
       ('supplyframe.app'),
       ('componentsearchengine.com');


insert into mta_status_text (short, support)
values ('Unknown',
        'Recipient''s MTA refused the email with a status error that we haven''t yet investigated.  Please contact ops@supplyframe.com for assistance.'),
       ('Sent successfully',
        'Email was sent successfully by the Supplyframe MTA and received by the recipient''s MTA without issue.  If the email is not seen by the recipient, the issue is guaranteed to be on the recipient''s side.  Typically it will be in the spam folder, otherwise recommend that the recipient contact their IT department.'),
       ('Recipient domain not found',
        'The recipient''s domain (the part to the right of the @) does not have a DNS entry and did not resolve to an IP address.  This means either the domain had a typo in it, the domain does not exist or the domain is not set up to receive email properly.'),
       ('Unable to connect to recipient''s MTA',
        'Supplyframe MTA attempted to contact the recipient''s MTA but we got a timeout or connection error.  The issue is likely to be a firewall at the recipient''s MTA or a network disruption or outage along the way to the recipient''s MTA.  Recommendation is to direct the recipient to their IT department to inquire about a firewall or network level issue that blocks email.'),
       ('Recipient mailbox unavailable',
        'The recipient''s MTA informed us that the inbox for the email address does not exist, is inactive or is not allowed to receive our email.  This typically happens when the email has a typo in the user name part (the left side of the @) or if the user was offboarded in the customer''s IT department.'),
       ('Recipient mailbox is full',
        'The recipient''s MTA informed us that the inbox for the email address was full and cannot accept any more emails.'),
       ('Bad reputation',
        'The recipient''s MTA blocked our attempt to send the email because they consider our MTA to have a bad reptuation.  This usually happens when public blacklists used by the recipient list our MTA in their database.  We get on those lists when people mark our emails as spam.  This should not be an issue for the domains we use in notificaiton emails because we do not share MTA nor domains with the marketing department.  The more likely issue that we''ve seen is that the IT department of the recipient has configurd a custom blacklist or is using an unusual blacklist that lists us for some reason.  We''ve seen situations where less common "aggressive" blacklists are used that are more likely to mark us with bad reptuation, notably in Germany.  We monitor for this issue and typically try to get ourselves off these blacklists, however it can take time and will not always succeed.  Please inquire to ops@supplyframe.com for an update on a particular email.')
;


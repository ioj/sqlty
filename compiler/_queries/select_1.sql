/*
  @name Whatever
  @param dupa -> ((user_id, book_name)...)
  @param spredzik -> (...)
  @paramAsStruct
  @notNullParams (notnull1, notnull2, nnul_3, dupa.user_id)
  @one
  @returnValueName whatever
*/
-- A to jest line comment
-- Inny line comment
-- Zupełnie inny line kąmet
--
--
-- Blebeo avxz

SELECT * from ng_private.verification_code
WHERE id = :id and isbns NOT IN :spredzik AND id != :id AND whore = :bitch;

/*
 @name other_one
 @paramAsStruct
 @param ids -> (...)
 @many
*/

-- OtherOne is a query which retrieves books
-- with given ids.

SELECT id, name, active FROM books
WHERE id IN :ids AND name LIKE '%blehblah%';
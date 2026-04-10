/*
  @name BulkInsertArrays
  @exec
  @param rows ((id, intTags, strTags, uuidRefs, flags)...)
*/
INSERT INTO e2e_arrays (id, int_tags, str_tags, uuid_refs, flags)
VALUES :rows;

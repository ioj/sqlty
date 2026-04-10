/* @name GetArrays @one */
SELECT int_tags, str_tags, uuid_refs, flags FROM e2e_arrays WHERE id = :id;

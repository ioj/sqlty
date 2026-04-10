/* @name FindByTags @many */
SELECT id, str_tags FROM e2e_arrays WHERE str_tags && :tags;

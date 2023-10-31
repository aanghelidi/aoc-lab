DROP SCHEMA IF EXISTS day03 CASCADE;


CREATE SCHEMA IF NOT EXISTS day03;


CREATE TABLE IF NOT EXISTS day03.inputs (s1 INTEGER NOT NULL CHECK (s1 > 0), s2 INTEGER NOT NULL CHECK (s2 > 0), s3 INTEGER NOT NULL CHECK (s3 > 0));

COPY day03.inputs
FROM 'input.txt' (
                  DELIMITER ',',
                            header FALSE);

WITH part_1_solution AS
  (SELECT sum(CASE
                  WHEN s1 + s2 <= s3
                       OR s1 + s3 <= s2
                       OR s2 + s3 <= s1 THEN 0
                  ELSE 1
              END)::int AS ans
   FROM day03.inputs),
     data_with_id AS
  (SELECT row_number() OVER (
                             ORDER BY
                               (SELECT 1))::INT AS id,
                            *
   FROM day03.inputs),
     data_with_group_id AS
  (SELECT ceil(row_number() OVER (
                                  ORDER BY id) / 3)::INT AS group_id,
          *
   FROM data_with_id),
     data_long AS (
                   FROM data_with_group_id unpivot (side_length
                                                    FOR triangle IN (s1, s2, s3))),
     part_2_possibles AS
  (SELECT triangle_id,
          side_length,
          CASE
              WHEN sum(side_length) OVER triangles - max(side_length) OVER triangles <= max(side_length) OVER triangles THEN 0
              ELSE 1
          END AS is_valid,
   FROM
     (SELECT group_id || '_' || triangle AS triangle_id,
             *
      FROM data_long
      ORDER BY triangle_id) WINDOW triangles AS (PARTITION BY triangle_id)),
     part_2_possibles_reduced AS
  (SELECT triangle_id,
          first(is_valid) AS is_valid,
   FROM part_2_possibles
   GROUP BY triangle_id),
     part_2_solution AS
  (SELECT sum(is_valid)::int AS ans
   FROM part_2_possibles_reduced)
SELECT p1.ans AS part_1,
       p2.ans AS part_2
FROM part_1_solution p1,
     part_2_solution p2;

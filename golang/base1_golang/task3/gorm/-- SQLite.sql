-- SQLite
select * from employees;

WITH ranked AS (
  SELECT *,
         DENSE_RANK() OVER (ORDER BY salary DESC) as rnk
  FROM employees
)
SELECT id, name, salary, department,rnk
FROM ranked 

SELECT * FROM `users` left join posts on posts.user_id=users.id left join comments on comments.post_id=posts.id WHERE users.name="alex"

insert into comments(id,content,post_id,user_id) values(38,	"张三对alex post1的评论1",	57,	43)
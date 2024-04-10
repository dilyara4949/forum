package repo

// user
const (
	createUserQuery     = `INSERT INTO users ( email, password, username) VALUES ($1, $2, $3);`
	getByEmailUserQuery = `SELECT username, email, password FROM users WHERE email = $1`
	// deleteByIDUserQuery = `UPDATE users SET is_deleted = true WHERE id = $1 and is_deleted=false returning id`
	getAllUsersQuery = `select username, email from users; `
	updateUserQuery  = `UPDATE users  SET username = $1 WHERE email = $2 returning email;`
	// setUserPassword     = `update users set password = $1 where id = $2 and is_deleted=false`
	// getAdmin            = `SELECT id, name, email, birthdate, language, phonenumber FROM users WHERE email = $1 and is_deleted=false and is_admin=true`
	// isAdmin             = `select is_admin from users where id = $1`
	getUserPassword = `select password from users where email=$1;`
)

// // category
const (
	createCategoryQuery = `insert into categories (name) values ($1);`
	getCategory         = `select name from categories where name = $1;`
	deleteCategory      = `delete from categories where name=$1;`
	getAllCategories    = `select name from categories;`
)

// post
const (
	createPostQuery = `insert into post (email, title, body, category) values ($1, $2, $3, $4) returning id;`
	getPost         = `select id, email, title, body, category from post where email = $1;`
	deletePost      = `delete from post where id=$1;`
	getAllPost      = `select id, email, title, body, category from post;`
	// IsOnePostLike   = `SELECT COUNT(*)  FROM like_post_user WHERE email=$1 and post_id = $2;`
	getOnePostLike  = `SELECT COUNT(*)  FROM like_post_user WHERE post_id = $1;`
)

// comment
const (
	createCommentQuery = `insert into comments (email, post_id, message) values ($1, $2, $3);`
	getComment         = `select email, post_id, message from comments where post_id = $1;`
	deleteComment      = `delete from comments where email=$1 and post_id = $2;`
	getAllComment      = `select email, post_id, message from comments;`
)
// like
const (
	createLikeQuery = `insert into like_post_user (email, post_id) values ($1, $2);`
	getLike         = `select email, post_id from like_post_user where email = $1 and post_id = $2;`
	deleteLike      = `delete from like_post_user where email=$1 and post_id = $2;`
)
// // genre
// const (
// 	createGenreQuery = `insert into genre (name, movieCount, link, fileid) values ($1, $2, $3, $4) returning id;`
// 	getGenre         = `select name, movieCount, link, fileid from genre where id = $1 ;`
// 	deleteGenre      = `delete from genre where id=$1 ;`
// 	getAllGenres     = `select id, name, movieCount, link, fileid from genre;`
// 	updateGenre      = `UPDATE genre SET name = $1, link=$2, fileid = $3 WHERE id = $4 returning id`
// )

// // poster
// const (
// 	createPoster  = `insert into poster (movieId, link, fileid) values ($1, $2, $3) returning id;`
// 	getPoster     = `select movieId, link, fileid from poster where id = $1;`
// 	deletePoster  = `delete from poster where id=$1';`
// 	getAllPosters = `select id, movieId, link, fileid from poster;`
// )

// // screenshot
// const (
// 	createScreenshot  = `insert into screenshot (movieId, link, fileid) values ($1, $2, $3) returning id;`
// 	getScreenshot     = `select movieId, link, fileid from screenshot where id = $1;`
// 	deleteScreenshot  = `delete from screenshot where id=$1;`
// )

// // commentage
// const (
// 	createVideo = `insert into Video (link, number, seasonid) values ($1, $2, $3) returning id;`
// 	getVideo         = `select link, number, seasonid from Video where id = $1;`
// 	deleteVideo      = `delete from Video where id=$1;`
// 	getAllVideos     = `select link, number, seasonid from Video;`
// 	updateVideo      = `UPDATE video SET link = $1, number=$2, seasonid = $3 WHERE id = $4;`
// )

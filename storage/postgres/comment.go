package postgres

import (
	"log"
	"time"

	c "bay_store/comment_service/genproto/comment"
)

func (r *CommentRepo) WriteComment(comment *c.CommentRequest) (*c.CommentResponse, error) {
	var res c.CommentResponse
	err := r.db.QueryRow(`
		insert into 
			comments(user_id, product_id, description) 
		values
			($1, $2, $3)
		returning 
			id, user_id, product_id, description, created_at`, comment.UserId, comment.ProductId, comment.Description).Scan(&res.Id, &res.UserId, &res.ProductId, &res.Description, &res.CreatedAt)

	if err != nil {
		log.Println("failed to write comment")
		return &c.CommentResponse{}, err
	}

	return &res, nil
}

func (r *CommentRepo) GetProductComments(req *c.IdRequest) (*c.Comments, error) {
	var res c.Comments
	rows, err := r.db.Query(`
		select 
			id, user_id, product_id, description, created_at
		from 
			comments 
		where
			product_id = $1 and deleted_at is null`, req.Id)

	if err != nil {
		log.Println("failed to get product comment")
		return &c.Comments{}, err
	}

	for rows.Next() {
		temp := c.CommentResponse{}

		err = rows.Scan(
			&temp.Id,
			&temp.UserId,
			&temp.ProductId,
			&temp.Description,
			&temp.CreatedAt,
		)
		if err != nil {
			log.Println("failed to scanning comment")
			return &c.Comments{}, err
		}

		res.Comments = append(res.Comments, &temp)
	}

	return &res, nil
}

func (r *CommentRepo) GetUserComments(req *c.IdRequest) (*c.Comments, error) {
	var res c.Comments
	rows, err := r.db.Query(`
		select 
			id, user_id, product_id, description, created_at
		from 
			comments 
		where
			user_id = $1 and deleted_at is null`, req.Id)

	if err != nil {
		log.Println("failed to get user comments")
		return &c.Comments{}, err
	}

	for rows.Next() {
		temp := c.CommentResponse{}

		err = rows.Scan(
			&temp.Id,
			&temp.UserId,
			&temp.ProductId,
			&temp.Description,
			&temp.CreatedAt,
		)
		if err != nil {
			log.Println("failed to scanning comment")
			return &c.Comments{}, err
		}

		res.Comments = append(res.Comments, &temp)
	}

	return &res, nil
}

func (r *CommentRepo) DeleteComment(id *c.IdRequest) (*c.CommentResponse, error) {
	temp := c.CommentResponse{}
	err := r.db.QueryRow(`
		update 
			comments
		set 
			deleted_at = $1 
		where 
			id = $2
		returning 
			id, user_id, product_id, description, created_at`, time.Now(), id.Id).Scan(&temp.Id, &temp.UserId, &temp.ProductId, &temp.Description, &temp.CreatedAt)

	if err != nil {
		log.Println("failed to delete comment")
		return &c.CommentResponse{}, err
	}

	return &temp, nil
}

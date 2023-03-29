package service

import (
	c "bay_store/comment_service/genproto/comment"
	p "bay_store/comment_service/genproto/product"
	u "bay_store/comment_service/genproto/user"
	"log"

	"bay_store/comment_service/pkg/logger"
	grpcclient "bay_store/comment_service/service/grpc_client"
	"bay_store/comment_service/storage"

	"context"

	"github.com/jmoiron/sqlx"
)

type CommentService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpcclient.Clients
}

func NewCommentService(db *sqlx.DB, log logger.Logger, client grpcclient.Clients) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
		Client:  client,
	}
}

func (s *CommentService) WriteComment(ctx context.Context, req *c.CommentRequest) (*c.CommentResponse, error) {
	res, err := s.storage.Comment().WriteComment(req)
	if err != nil {
		log.Println("failed to write comment: ", err)
		return &c.CommentResponse{}, err
	}

	product, err := s.Client.Product().GetProductById(context.Background(), &p.IdRequest{Id: res.ProductId})
	if err != nil {
		log.Println("failed to get product: ", err)
		return &c.CommentResponse{}, err
	}
	res.ProductName = product.Name

	user, err := s.Client.User().GetUserById(context.Background(), &u.IdRequest{Id: res.UserId})
	if err != nil {
		log.Println("failed to get user: ", err)
		return &c.CommentResponse{}, err
	}
	res.UserName = user.FirstName + " " + user.LastName

	return res, nil
}

func (s *CommentService) GetProductComments(ctx context.Context, req *c.IdRequest) (*c.Comments, error) {
	res, err := s.storage.Comment().GetProductComments(req)
	if err != nil {
		log.Println("failed to get product comments: ", err)
		return &c.Comments{}, err
	}

	for _, val := range res.Comments {
		product, err := s.Client.Product().GetProductById(context.Background(), &p.IdRequest{Id: val.ProductId})
		if err != nil {
			log.Println("failed to get product: ", err)
			return &c.Comments{}, err
		}
		val.ProductName = product.Name

		user, err := s.Client.User().GetUserById(context.Background(), &u.IdRequest{Id: val.UserId})
		if err != nil {
			log.Println("failed to get user: ", err)
			return &c.Comments{}, err
		}
		val.UserName = user.FirstName + " " + user.LastName
	}

	return res, nil
}

func (s *CommentService) GetUserComments(ctx context.Context, req *c.IdRequest) (*c.Comments, error) {
	res, err := s.storage.Comment().GetUserComments(req)
	if err != nil {
		log.Println("failed to get user comments: ", err)
		return &c.Comments{}, err
	}

	for _, val := range res.Comments {
		product, err := s.Client.Product().GetProductById(context.Background(), &p.IdRequest{Id: val.ProductId})
		if err != nil {
			log.Println("failed to get product: ", err)
			return &c.Comments{}, err
		}
		val.ProductName = product.Name

		user, err := s.Client.User().GetUserById(context.Background(), &u.IdRequest{Id: val.UserId})
		if err != nil {
			log.Println("failed to get user: ", err)
			return &c.Comments{}, err
		}
		val.UserName = user.FirstName + " " + user.LastName
	}

	return res, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *c.IdRequest) (*c.CommentResponse, error) {
	res, err := s.storage.Comment().DeleteComment(req)
	if err != nil {
		log.Println("failed to get product by id: ", err)
		return &c.CommentResponse{}, err
	}

	product, err := s.Client.Product().GetProductById(context.Background(), &p.IdRequest{Id: res.ProductId})
	if err != nil {
		log.Println("failed to get product: ", err)
		return &c.CommentResponse{}, err
	}
	res.ProductName = product.Name

	user, err := s.Client.User().GetUserById(context.Background(), &u.IdRequest{Id: res.UserId})
	if err != nil {
		log.Println("failed to get user: ", err)
		return &c.CommentResponse{}, err
	}
	res.UserName = user.FirstName + " " + user.LastName

	return res, nil
}

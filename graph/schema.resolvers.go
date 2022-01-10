package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grapqh-relation/graph/generated"
	"cesargdd/grapqh-relation/graph/model"
	"cesargdd/grapqh-relation/pg"
	"context"
	"fmt"
)

var conn = pg.Connect()
var db = pg.New(conn)

func (r *commentResolver) Likes(ctx context.Context, obj *pg.Comment) ([]pg.Commentlike, error) {
	commentId, err := db.GetComment(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	commentLikes, err := db.LikesComment(context.Background(), commentId.ID)
	if err != nil {
		fmt.Println(err)
	}
	return commentLikes, nil
}

func (r *commentlikeResolver) User(ctx context.Context, obj *pg.Commentlike) (*pg.User, error) {
	commentLikeId, err := db.GetCommentLike(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	user, err := db.GetUserID(context.Background(), commentLikeId.Userid)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*pg.User, error) {
	user, err := db.CreateUser(context.Background(), pg.CreateUserParams{
		Username: input.Username,
		Bio:      *input.Bio,
		Avatar:   *input.Avatar,
		Email:    *input.Email,
		Password: input.Password,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.EditUser) (*pg.User, error) {
	user, err := db.UpdateUser(context.Background(), pg.UpdateUserParams{
		ID:     int32(input.ID),
		Bio:    *input.Bio,
		Avatar: *input.Avatar,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *int) (*pg.User, error) {
	deleteUser, err := db.DeleteUser(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteUser, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*pg.Post, error) {
	post, err := db.CreatePost(context.Background(), pg.CreatePostParams{
		Url:     input.URL,
		Caption: *input.Caption,
		Userid:  int32(input.UserID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        post.ID,
		Url:       post.Url,
		Caption:   post.Caption,
		CreatedAt: post.CreatedAt,
		Userid:    post.Userid,
	}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input *model.EditPost) (*pg.Post, error) {
	updatedPost, err := db.UpdatePost(context.Background(), pg.UpdatePostParams{
		ID:      int32(input.ID),
		Url:     *input.URL,
		Caption: *input.Caption,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        updatedPost.ID,
		Url:       updatedPost.Url,
		Caption:   updatedPost.Caption,
		UpdatedAt: updatedPost.UpdatedAt,
		CreatedAt: updatedPost.CreatedAt,
		Userid:    updatedPost.Userid,
	}, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *int) (*pg.Post, error) {
	deletePost, err := db.DeletePost(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deletePost, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*pg.Comment, error) {
	comment, err := db.CreateComment(context.Background(), pg.CreateCommentParams{
		Contents: input.Contents,
		Userid:   int32(input.UserID),
		Postid:   int32(input.PostID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		Userid:    comment.Userid,
		Postid:    comment.Postid,
	}, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, input *model.EditComment) (*pg.Comment, error) {
	updateComment, err := db.UpdateComment(context.Background(), pg.UpdateCommentParams{
		ID:       int32(input.ID),
		Contents: input.Contents,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        updateComment.ID,
		CreatedAt: updateComment.CreatedAt,
		UpdatedAt: updateComment.UpdatedAt,
		Contents:  updateComment.Contents,
		Userid:    updateComment.Userid,
		Postid:    updateComment.Postid,
	}, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, input *int) (*pg.Comment, error) {
	deleteComment, err := db.DeleteComment(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteComment, nil
}

func (r *mutationResolver) CreatePostLike(ctx context.Context, input *model.NewPostLike) (*pg.Postlike, error) {
	like, err := db.CreatePostLike(context.Background(), pg.CreatePostLikeParams{
		Userid: int32(input.UserID),
		Postid: int32(input.PostID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Postlike{
		ID:        like.ID,
		CreatedAt: like.CreatedAt,
		Userid:    like.Userid,
		Postid:    like.Postid,
	}, nil
}

func (r *mutationResolver) DeletePostLike(ctx context.Context, input *int) (*pg.Postlike, error) {
	deleteLike, err := db.DeletePostLike(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteLike, nil
}

func (r *mutationResolver) CreateCommentLike(ctx context.Context, input *model.NewCommentLike) (*pg.Commentlike, error) {
	commentLike, err := db.CreateCommentLike(context.Background(), pg.CreateCommentLikeParams{
		Userid:    int32(input.UserID),
		Commentid: int32(input.CommentID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Commentlike{
		ID:        commentLike.ID,
		CreatedAt: commentLike.CreatedAt,
		Userid:    commentLike.Userid,
		Commentid: commentLike.Commentid,
	}, nil
}

func (r *mutationResolver) DeleteCommentLike(ctx context.Context, input *int) (*pg.Commentlike, error) {
	delComLike, err := db.DeleteCommentLike(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &delComLike, nil
}

func (r *mutationResolver) CreateHashtag(ctx context.Context, input *model.NewHashtag) (*pg.Hashtag, error) {
	createHashtag, err := db.CreateHashtag(context.Background(), input.Title)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        createHashtag.ID,
		CreatedAt: createHashtag.CreatedAt,
		Title:     createHashtag.Title,
	}, nil
}

func (r *mutationResolver) UpdateHashtag(ctx context.Context, input *model.EditHashtag) (*pg.Hashtag, error) {
	updateHashtag, err := db.UpdateHashtag(context.Background(), pg.UpdateHashtagParams{
		ID:    int32(input.ID),
		Title: input.Title,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        updateHashtag.ID,
		CreatedAt: updateHashtag.CreatedAt,
		Title:     updateHashtag.Title,
	}, nil
}

func (r *mutationResolver) DeleteHashtag(ctx context.Context, input *int) (*pg.Hashtag, error) {
	deleteHashtag, err := db.DeleteHashtag(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteHashtag, nil
}

func (r *mutationResolver) CreateFollower(ctx context.Context, input *model.NewFollower) (*pg.Follower, error) {
	follower, err := db.CreateFollower(context.Background(), pg.CreateFollowerParams{
		Leaderid:   int32(input.LeaderID),
		Followerid: int32(input.FollowerID),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Follower{
		ID:         follower.ID,
		Leaderid:   follower.Leaderid,
		Followerid: follower.Followerid,
		CreatedAt:  follower.CreatedAt,
	}, nil
}

func (r *mutationResolver) DeleteFollower(ctx context.Context, input *int) (*pg.Follower, error) {
	deleteFollower, err := db.DeleteFollower(context.Background(), int32(*input))
	if err != nil {
		fmt.Println(err)
	}
	return &deleteFollower, nil
}

func (r *postResolver) Comments(ctx context.Context, obj *pg.Post) ([]pg.Comment, error) {
	postId, err := db.GetPost(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	commentsPost, err := db.ListCommentsPost(context.Background(), postId.ID)
	if err != nil {
		fmt.Println(err)
	}
	return commentsPost, nil
}

func (r *postResolver) Likes(ctx context.Context, obj *pg.Post) ([]pg.Postlike, error) {
	postId, err := db.GetPost(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	postLikes, err := db.LikesPost(context.Background(), postId.ID)
	if err != nil {
		fmt.Println(err)
	}
	return postLikes, nil
}

func (r *postlikeResolver) User(ctx context.Context, obj *pg.Postlike) (*pg.User, error) {
	postId, err := db.GetPostLike(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	user, err := db.GetUserID(context.Background(), postId.Postid)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, username string) (*pg.User, error) {
	user, err := db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.User{
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]pg.User, error) {
	users, err := db.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return users, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*pg.Post, error) {
	getPost, err := db.GetPost(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Post{
		ID:        getPost.ID,
		Url:       getPost.Url,
		Caption:   getPost.Caption,
		UpdatedAt: getPost.UpdatedAt,
		CreatedAt: getPost.CreatedAt,
		Userid:    getPost.Userid,
	}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]pg.Post, error) {
	posts, err := db.ListPosts(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

func (r *queryResolver) Comment(ctx context.Context, id int) (*pg.Comment, error) {
	comment, err := db.GetComment(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Comment{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		Userid:    comment.Userid,
		Postid:    comment.Postid,
	}, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]pg.Comment, error) {
	comments, err := db.ListComments(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return comments, nil
}

func (r *queryResolver) PostLike(ctx context.Context, id int) (*pg.Postlike, error) {
	postLike, err := db.GetPostLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Postlike{
		ID:        postLike.ID,
		CreatedAt: postLike.CreatedAt,
		Userid:    postLike.Userid,
		Postid:    postLike.Postid,
	}, nil
}

func (r *queryResolver) PostLikes(ctx context.Context) ([]pg.Postlike, error) {
	postLikes, err := db.ListPostLikes(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return postLikes, nil
}

func (r *queryResolver) CommentLike(ctx context.Context, id int) (*pg.Commentlike, error) {
	commentLike, err := db.GetCommentLike(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Commentlike{
		ID:        commentLike.ID,
		CreatedAt: commentLike.CreatedAt,
		Userid:    commentLike.Userid,
		Commentid: commentLike.Commentid,
	}, nil
}

func (r *queryResolver) CommentLikes(ctx context.Context) ([]pg.Commentlike, error) {
	commentLikes, err := db.ListCommentLikes(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return commentLikes, nil
}

func (r *queryResolver) Hashtag(ctx context.Context, title string) (*pg.Hashtag, error) {
	hashtag, err := db.GetHashtag(context.Background(), title)
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Hashtag{
		ID:        hashtag.ID,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *queryResolver) Hashtags(ctx context.Context) ([]pg.Hashtag, error) {
	hashtags, err := db.ListHashtags(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return hashtags, nil
}

func (r *queryResolver) Follower(ctx context.Context, id int) (*pg.Follower, error) {
	follower, err := db.GetFollower(context.Background(), int32(id))
	if err != nil {
		fmt.Println(err)
	}
	return &pg.Follower{
		ID:         follower.ID,
		Leaderid:   follower.Leaderid,
		Followerid: follower.Followerid,
		CreatedAt:  follower.CreatedAt,
	}, nil
}

func (r *queryResolver) Followers(ctx context.Context) ([]pg.Follower, error) {
	followers, err := db.ListFollowers(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return followers, nil
}

func (r *userResolver) Posts(ctx context.Context, obj *pg.User) ([]pg.Post, error) {
	user, err := db.GetUserID(context.Background(), obj.ID)
	if err != nil {
		fmt.Println(err)
	}
	posts, err := db.ListPostsUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Commentlike returns generated.CommentlikeResolver implementation.
func (r *Resolver) Commentlike() generated.CommentlikeResolver { return &commentlikeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Postlike returns generated.PostlikeResolver implementation.
func (r *Resolver) Postlike() generated.PostlikeResolver { return &postlikeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type commentlikeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type postlikeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

func (r *postResolver) Comment(ctx context.Context, obj *pg.Post) (*pg.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

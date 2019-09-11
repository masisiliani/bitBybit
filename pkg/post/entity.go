package post

type Post struct {
	ID          int
	User        string
	Description string
}

type PostController struct {
	Repository Repository
}

func (pc *PostController) GetPostByID(id int) (Post, error){
	post, err := pc.Repository.FindPost(id)
	return post, err
}

func  (pc *PostController)  GetPosts(username string) ([]Post, error){
	posts, err := pc.Repository.FindPosts(username)
	return posts, err
}

func  (pc *PostController) InsertPost(p Post){

}

func  (pc *PostController) DeletePost(id int) error{

}

func  (pc *PostController) UpdatePost(p Post){

}
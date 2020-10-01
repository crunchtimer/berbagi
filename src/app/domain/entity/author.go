package entity

type Author struct {
	DisplayName string
	UserName    string
	ImageUrl    string
}

func NewAuthor(displayName, userName, ImageUrl string) (author *Author, err error) {
	author = new(Author)
	author.DisplayName = displayName
	author.UserName = userName
	author.ImageUrl = ImageUrl
	return author, nil
}

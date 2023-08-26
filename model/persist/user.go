package persist

type User struct {
	UserID string

	RemainCount int64

	InviteCode string

	UserName string

	IsDonate int64
}

func (u *User) Donated() bool {
	return u.IsDonate == 1
}

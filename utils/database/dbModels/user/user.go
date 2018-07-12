package user

type UserAccount struct {
	id        int    `db:"id"`
	wechat_id string `db:"w_id"`
}

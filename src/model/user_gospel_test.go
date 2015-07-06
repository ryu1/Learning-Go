package model

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestUser(t *testing.T) {

	Describe(t, "NewUser", func() {

		Context("フルネーム（hoge aaa）を指定した場合", func() {
			user := NewUser("hoge aaa")

			It("FirstName should be hoge", func() {
				Expect(user.FirstName).To(Equal, "hoge")
			})

			It("LastName should be aaa", func() {
				Expect(user.LastName).To(Equal, "aaa")
			})

		})

		Context("空文字を渡した場合", func() {
			user := NewUser("")

			It("FirstName should be empty string", func() {
				Expect(user.FirstName).To(Equal, "")
			})

			It("LastName should be empty string", func() {
				Expect(user.LastName).To(Equal, "")
			})

		})

	})
	Describe(t, "Divisions", func() {
		user := NewUser("")

		It("default divisions is empty slice", func() {
			Expect(len(user.Divisions)).To(Equal, 0)
		})

	})

}

package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerNameCannotBeBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	// กรณี Name เป็นค่าว่าง ""
	customer := Customer{
		Name:       "", // ผิดตรงนี้
		Email:      "test@example.com",
		CustomerID: "L1234567",
	}

	// ตรวจสอบ
	ok, err := govalidator.ValidateStruct(customer)

	// คาดหวังว่า ok ต้องเป็น false และ err ต้องไม่เป็น nil
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	// คาดหวังข้อความ Error ตรงกับที่เขียนไว้ใน struct tag
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}
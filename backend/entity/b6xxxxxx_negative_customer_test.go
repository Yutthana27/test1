package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerIDInvalidFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	// กรณี CustomerID ผิดรูปแบบ (เช่น ขึ้นต้นด้วย K ซึ่งไม่ได้อยู่ใน L, M, H)
	customer := Customer{
		Name:       "Yutthana",
		Email:      "test@example.com",
		CustomerID: "K1234567", // ผิดตรงนี้ (ขึ้นต้นด้วย K)
	}

	// ตรวจสอบ
	ok, err := govalidator.ValidateStruct(customer)

	// คาดหวังว่า ok ต้องเป็น false และ err ต้องไม่เป็น nil
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	// คาดหวังข้อความ Error ตรงกับที่เขียนไว้ใน struct tag
	g.Expect(err.Error()).To(Equal("CustomerID invalid format"))
}
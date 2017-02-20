/*
	Created On:8 Feb 2017
	Author:Nilav

	Implementing Test Cases for testing different payroll management operations
 */

package test_test

import (
	"github.com/structure/futsalStruct"
	"../programs/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/json"
	"fmt"
	"time"
)

var _ = Describe("PayrollTest", func() {

	var (
		header futsalStruct.Header
		users futsalStruct.User
		payroll futsalStruct.Payroll
		jsonData string
	)

	BeforeEach(func() {
		header = futsalStruct.Header{"Add Payroll", "ADD"}
		users = futsalStruct.User{"1", "Nilav", "something@gmail.com", "Tech", "Khatiwada", "password", "9843241699"}
		payroll = futsalStruct.Payroll{"ID123", users, "200", time.Now().UTC()}

		b, _ := json.Marshal(payroll)
		jsonData = string(b)
	})

	Describe("Testing the Payroll", func() {
		Context("Testing the adding of payroll informatioin", func() {

			It("Should return successful", func() {
				fmt.Println("There is something   ", jsonData)
				Expect(service.AddPayrollInfo(payroll)).To(Equal("Successfull"))
			})

		})

		Context("Getting Payroll By Payroll Id", func() {

			It("Should return successfully signed up", func() {
				fmt.Println("There is something   ", jsonData)
				Expect(service.GetPayrollInfoById(payroll.PayrollId)).To(Equal("Successfull"))
			})

			Context("Updating Payroll", func() {

				It("Should return successfully signed up", func() {
					fmt.Println("There is something   ", jsonData)
					Expect(service.UpdatePayrollById(payroll.PayrollId, payroll)).To(Equal("Successfull"))
				})

			})

			Context("Deleting User", func() {

				It("Should return successfully signed up", func() {
					fmt.Println("There is something   ", jsonData)
					Expect(service.DeletePayrollById(payroll.PayrollId)).To(Equal("Successfull"))
				})

			})
		})
	})
})

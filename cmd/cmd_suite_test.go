//
// Copyright (c) 2017 Dell Inc. or its subsidiaries.  All Rights Reserved.
// Dell EMC Confidential/Proprietary Information
//
//

package cmd_test

import (
	"flag"
	"time"

	"github.com/dellemc-symphony/workflow-cli/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/gin-gonic/gin.v1"

	"testing"
)

var https bool

func init() {
	flag.BoolVar(&https, "https", false, "Set 'true' to enable HTTPS for mock REST endpoint")
}

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmd Suite")
}

var router *gin.Engine

var _ = BeforeSuite(func() {
	mock.CreateMock(https)
	time.Sleep(25 * time.Millisecond)
})

var _ = AfterSuite(func() {
	mock.StopMock()
})

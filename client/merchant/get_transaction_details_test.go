package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionDetails(t *testing.T) {
	assert := assert.New(t)

	cli := testNewDefault()
	// error
	svc := &GetTransactionDetails{
		TransactionID: "XXXXXXXXXXXXX",
	}
	v, err := svc.Do(cli)

	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10004", v.ErrorCode)

}

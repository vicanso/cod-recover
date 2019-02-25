// Copyright 2018 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recover

import (
	"fmt"
	"net/http"

	"github.com/vicanso/hes"

	"github.com/vicanso/cod"
)

const (
	// ErrCategoryRecover recover error category
	ErrCategoryRecover = "cod-recover"
)

// New new recover
func New() cod.Handler {
	return func(c *cod.Context) error {
		defer func() {
			// 此recover只是示例，在实际使用中，
			// 需要针对实际需求调整，如对于每个recover增加邮件通知等
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					he := hes.NewWithErrorStatusCode(fmt.Errorf("%v", r), http.StatusInternalServerError)
					he.Category = ErrCategoryRecover
					err = he
				}
				c.Cod(nil).EmitError(c, err)
				// 如果已直接对Response写入数据，则将 Committed设置为 true
				c.Committed = true
				resp := c.Response
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
			}
		}()
		return c.Next()
	}
}

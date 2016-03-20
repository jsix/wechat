<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package poi

import (
	"github.com/chanxuehong/wechat/mp"
)

// 查询门店信息.
func (clt *Client) GetWxCategory() (categoryList []string, err error) {
	var result struct {
		mp.Error
		CategoryList []string `json:"category_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/api_getwxcategory?access_token="
	if err = ((*mp.Client)(clt)).GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	categoryList = result.CategoryList
=======
package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// CategoryList 获取门店类目表.
func CategoryList(clt *core.Client) (list []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/api_getwxcategory?access_token="

	var result struct {
		core.Error
		CategoryList []string `json:"category_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.CategoryList
>>>>>>> github/v2
	return
}

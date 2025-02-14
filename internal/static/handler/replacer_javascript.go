//
// Copyright (C) 2023 Quan Chen <chenquan_act@163.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handler

import (
	"github.com/terasum/medict/pkg/model"
	"regexp"
	"strings"
)

var _ Replacer = &ReplacerJs{}

var JS_REG *regexp.Regexp

func init() {
	var err error
	JS_REG, err = regexp.Compile(`src=\"(\S+\.js)\"`)
	if err != nil {
		panic(err)
	}
}

type ReplacerJs struct {
	mdict *model.DictionaryItem
}

func (r *ReplacerJs) SetDictContext(item *model.DictionaryItem) {
	r.mdict = item
}

func (r *ReplacerJs) Replace(dictId string, entry *model.MdictKeyWordIndex, html string) (*model.MdictKeyWordIndex, string) {
	if html == "" || dictId == "" {
		return entry, html
	}

	newhtml := html
	matchedGroup := JS_REG.FindAllStringSubmatch(html, -1)
	for _, matched := range matchedGroup {
		if len(matched) != 2 {
			continue
		}
		oldStr := matched[1]
		newStr := oldStr + "?dict_id=" + dictId + "&d=0"
		newhtml = strings.ReplaceAll(newhtml, oldStr, newStr)
	}

	return entry, newhtml
}

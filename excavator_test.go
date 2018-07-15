package excavator_test

import (
	"log"
	"strings"
	"testing"

	"github.com/godcong/excavator"
	"github.com/godcong/excavator/db"
)

func TestRadical_Iterator(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	root := excavator.NewRoot("http://tool.httpcn.com", "/KangXi/BuShou.html")
	root.Self()
	radical := root.SelfRadical("牛")
	log.Println(radical)
	if radical == nil {
		return
	}
	//log.Println(radical.SelfCharacter("乊"))
	radical.Iterator(func(character *excavator.Character) error {
		if character.Character == "犪" {
			log.Println(character)
			db.DB().Insert(character)
		}

		return nil
	})

}

func TestRadical_Add(t *testing.T) {
	text := "汉字五行：土　是否为常用字：否"
	s := strings.SplitAfter(text, "：")
	log.Println(s, len(s))

}

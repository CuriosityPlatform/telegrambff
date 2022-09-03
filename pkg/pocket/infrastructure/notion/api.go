package notion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jomei/notionapi"
)

func NewApi(secretKey string) *Api {
	return &Api{
		client: notionapi.NewClient(notionapi.Token(secretKey)),
	}
}

type Api struct {
	client *notionapi.Client
}

func (api *Api) GetPage(pageID string) {
	//page, err := api.client.Page.Get(context.Background(), notionapi.PageID(pageID))
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("PAGE", page)

	//page, err := api.client.Database.Get(context.Background(), notionapi.DatabaseID(pageID))
	//if err != nil {
	//	panic(err)
	//}
	//
	//marshal, err := json.Marshal(page)
	//fmt.Println("PAGE", string(marshal))

	page, err := api.client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(pageID),
		},
		Properties: map[string]notionapi.Property{
			"Name": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{{
					Type: notionapi.ObjectTypeText,
					Text: notionapi.Text{
						Content: "Name",
					},
				}},
			},
			"URL": notionapi.URLProperty{
				Type: notionapi.PropertyTypeURL,
				URL:  "https://github.com/go-telegram-bot-api/telegram-bot-api",
			},
		},
		Cover: &notionapi.Image{
			Type: notionapi.FileTypeExternal,
			External: &notionapi.FileObject{
				//URL: "https://images.unsplash.com/opengraph/1x1.png?auto=format&fit=crop&w=1200&h=630&q=60&mark-w=64&mark-align=top%2Cleft&mark-pad=50&blend-w=1&mark=https%3A%2F%2Fimages.unsplash.com%2Fopengraph%2Flogo.png&blend=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1661180359798-6f263f523ac2%3Fcrop%3Dfaces%252Cedges%26cs%3Dtinysrgb%26fit%3Dcrop%26fm%3Djpg%26ixid%3DMnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjYyMjIwNzY4%26ixlib%3Drb-1.2.1%26q%3D60%26w%3D1200%26auto%3Dformat%26h%3D630%26mark-w%3D424%26mark-align%3Dmiddle%252Ccenter%26blend-mode%3Dnormal%26blend-alpha%3D10%26mark%3Dhttps%253A%252F%252Fimages.unsplash.com%252Fopengraph%252Fwordmark.png%26blend%3D000000",
				URL: "https://opengraph.githubassets.com/0c83e923eb27812d76e36438da64efa213f94673270685262cda5612a17585f1/go-telegram-bot-api/telegram-bot-api",
			},
		},
		//Children: []notionapi.Block{
		//	notionapi.ImageBlock{
		//		BasicBlock: notionapi.BasicBlock{
		//			Object: notionapi.ObjectTypeBlock,
		//			Type:   notionapi.BlockTypeImage,
		//		},
		//		Image: notionapi.Image{
		//			Type: notionapi.FileTypeExternal,
		//			External: &notionapi.FileObject{
		//				//URL: "https://images.unsplash.com/opengraph/1x1.png?auto=format&fit=crop&w=1200&h=630&q=60&mark-w=64&mark-align=top%2Cleft&mark-pad=50&blend-w=1&mark=https%3A%2F%2Fimages.unsplash.com%2Fopengraph%2Flogo.png&blend=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1661180359798-6f263f523ac2%3Fcrop%3Dfaces%252Cedges%26cs%3Dtinysrgb%26fit%3Dcrop%26fm%3Djpg%26ixid%3DMnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjYyMjIwNzY4%26ixlib%3Drb-1.2.1%26q%3D60%26w%3D1200%26auto%3Dformat%26h%3D630%26mark-w%3D424%26mark-align%3Dmiddle%252Ccenter%26blend-mode%3Dnormal%26blend-alpha%3D10%26mark%3Dhttps%253A%252F%252Fimages.unsplash.com%252Fopengraph%252Fwordmark.png%26blend%3D000000",
		//				URL: "https://opengraph.githubassets.com/0c83e923eb27812d76e36438da64efa213f94673270685262cda5612a17585f1/go-telegram-bot-api/telegram-bot-api",
		//			},
		//		},
		//	},
		//},
	})
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(page)
	fmt.Println("PAGE", string(marshal))
}

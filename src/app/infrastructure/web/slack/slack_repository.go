package slack_webhook

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	"net/http"
	"os"
	"time"
)

type SlackRepository struct {
}

func NewSlackRepository() *SlackRepository {
	repo := new(SlackRepository)
	return repo
}

func (repo *SlackRepository) Send(tallyCount []value_object.TallyCountItem) (err error) {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	now := time.Now()
	nowStr := now.Format("2006年1月2日")
	past := now.AddDate(0, 0, -7)
	pastStr := past.Format("2006年1月2日")

	var body string
	switch len(tallyCount) {
	case 1:
		body = fmt.Sprintf(`
{                                                                                 
    "blocks": [
    	{
    		"type": "section",
    		"text": {
    			"type": "mrkdwn",
    			"text": "%s〜%s"
    		}
    	},
        {
                "type": "section",
                "block_id": "1",
                "text": {
                        "type": "mrkdwn",
                  		"text": " >>> :first_place_medal: 1位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        }
    ]
}
`,
			pastStr,
			nowStr,
			tallyCount[0].Author.DisplayName,
			tallyCount[0].Author.UserName,
			tallyCount[0].LikeCount,
			tallyCount[0].Author.ImageUrl,
			tallyCount[0].Author.UserName)

	case 2:
		body = fmt.Sprintf(`
{                                                                                 
    "blocks": [
    	{
    		"type": "section",
    		"text": {
    			"type": "mrkdwn",
    			"text": "%s〜%s"
    		}
    	},
        {
                "type": "section",
                "block_id": "1",
                "text": {
                        "type": "mrkdwn",
                  		"text": " >>> :first_place_medal: 1位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        },
        {
                "type": "section",
                "block_id": "2",
                "text": {
                        "type": "mrkdwn",
                  		"text": " >>> :second_place_medal: 2位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        },
    ]
}
`,
			pastStr,
			nowStr,
			tallyCount[0].Author.DisplayName,
			tallyCount[0].Author.UserName,
			tallyCount[0].LikeCount,
			tallyCount[0].Author.ImageUrl,
			tallyCount[0].Author.UserName,
			tallyCount[1].Author.DisplayName,
			tallyCount[1].Author.UserName,
			tallyCount[1].LikeCount,
			tallyCount[1].Author.ImageUrl,
			tallyCount[1].Author.UserName,
		)

	default:
		body = fmt.Sprintf(`
{                                                                                 
    "blocks": [
    	{
    		"type": "section",
    		"text": {
    			"type": "mrkdwn",
    			"text": "%s〜%s"
    		}
    	},
        {
                "type": "section",
                "block_id": "1",
                "text": {
                        "type": "mrkdwn",
                  		"text": " >>> :first_place_medal: 1位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        },
        {
                "type": "section",
                "block_id": "2",
                "text": {
                        "type": "mrkdwn",
                  		"text": " >>> :second_place_medal: 2位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        },
        {
                "type": "section",
                "block_id": "3",
                "text": {
                        "type": "mrkdwn",
                  		"text": ">>>  :third_place_medal: 3位 \n %s \n %s \n いいね数 %d"
                },
                "accessory": {
                        "type": "image",
                        "image_url": "%s",
                        "alt_text": "%s"
                }
        },
    ]
}
`,
			pastStr,
			nowStr,
			tallyCount[0].Author.DisplayName,
			tallyCount[0].Author.UserName,
			tallyCount[0].LikeCount,
			tallyCount[0].Author.ImageUrl,
			tallyCount[0].Author.UserName,
			tallyCount[1].Author.DisplayName,
			tallyCount[1].Author.UserName,
			tallyCount[1].LikeCount,
			tallyCount[1].Author.ImageUrl,
			tallyCount[1].Author.UserName,
			tallyCount[2].Author.DisplayName,
			tallyCount[2].Author.UserName,
			tallyCount[2].LikeCount,
			tallyCount[2].Author.ImageUrl,
			tallyCount[2].Author.UserName,
		)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(res.Status)
		return
	}
	defer res.Body.Close()
	return
}

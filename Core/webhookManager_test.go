package Core

import "testing"

func TestWebhookSend(t *testing.T) {
	var manga = DbMangaEntry{
		Did:          1,
		Dmanga:       "Test",
		DlastChapter: 1,
		Dmonitoring:  true,
		DchapterLink: "https://www.google.com",
		Didentifier:  "test",
	}
	t.Run(t.Name(), func(t *testing.T) {
		WebhookSend(manga)
	})

	//type args struct {
	//	manga DbMangaEntry
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		WebhookSend(tt.args.manga)
	//	})
	//}
}

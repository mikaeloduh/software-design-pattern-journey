package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelSubscriber(t *testing.T) {
	var writer bytes.Buffer
	ch1 := FakeNewChannel("水球軟體學院", &writer)
	ch2 := FakeNewChannel("PewDiePie", &writer)
	waterBall := NewSubscriber("水球")
	fireBall := NewSubscriber("火球")

	unsubscribe := func(c *Channel, v *Video, this *Subscriber) {
		if v.Length <= 60 {
			c.Unsubscribe(this)
		}
	}
	likeVideo := func(_ *Channel, v *Video, this *Subscriber) {
		if v.Length >= 180 {
			v.Like(this)
		}
	}
	waterBall.SetNotify(likeVideo)
	fireBall.SetNotify(unsubscribe)

	t.Run("水球訂閱 PewDiePie 和 水球軟體學院", func(t *testing.T) {
		ch1.Subscribe(waterBall)

		assert.Equal(t, "水球 訂閱了 水球軟體學院。", writer.String())
		writer.Reset()

		ch2.Subscribe(waterBall)

		assert.Equal(t, "水球 訂閱了 PewDiePie。", writer.String())
		writer.Reset()
	})

	t.Run("火球訂閱 PewDiePie 和 水球軟體學院", func(t *testing.T) {
		ch1.Subscribe(fireBall)

		assert.Equal(t, "火球 訂閱了 水球軟體學院。", writer.String())
		writer.Reset()

		ch2.Subscribe(fireBall)

		assert.Equal(t, "火球 訂閱了 PewDiePie。", writer.String())
		writer.Reset()
	})

	t.Run("水球軟體學院上傳一部影片：標題：”C1M1S2”、敘述：”這個世界正是物件導向的呢！”、影片長度：4 分鐘。", func(t *testing.T) {
		ch1.Upload(&Video{
			Title:       "C1M1S2",
			Description: "這個世界正是物件導向的呢！",
			Length:      240,
			Writer:      &writer,
		})

		assert.Contains(t, writer.String(), "頻道 水球軟體學院 上架了一則新影片 \"C1M1S2\"。")
		assert.Contains(t, writer.String(), "水球 對影片 \"C1M1S2\" 按讚。")
		writer.Reset()
	})

	t.Run("PewDiePie 上傳一部影片：標題：”Hello guys”、敘述：”Clickbait”、影片長度：30 秒。", func(t *testing.T) {
		ch2.Upload(&Video{
			Title:       "Hello guys",
			Description: "Clickbait",
			Length:      30,
			Writer:      &writer,
		})

		assert.Contains(t, writer.String(), "頻道 PewDiePie 上架了一則新影片 \"Hello guys\"。")
		assert.Contains(t, writer.String(), "火球 解除訂閱了 PewDiePie。")
		writer.Reset()
	})

	t.Run("水球軟體學院上傳一部影片：標題：”C1M1S3”、敘述：”物件 vs. 類別”、影片長度：1 分鐘。", func(t *testing.T) {
		ch1.Upload(&Video{
			Title:       "C1M1S3",
			Description: "物件 vs. 類別",
			Length:      60,
			Writer:      &writer,
		})

		assert.Contains(t, writer.String(), "頻道 水球軟體學院 上架了一則新影片 \"C1M1S3\"。")
		assert.Contains(t, writer.String(), "火球 解除訂閱了 水球軟體學院。")
		writer.Reset()
	})

	t.Run("PewDiePie 上傳一部影片：標題：”Minecraft”、敘述：”Let’s play Minecraft”、影片長度：30 分鐘。", func(t *testing.T) {
		ch2.Upload(&Video{
			Title:       "Minecraft",
			Description: "Let’s play Minecraft",
			Length:      180,
			Writer:      &writer,
		})

		assert.Contains(t, writer.String(), "頻道 PewDiePie 上架了一則新影片 \"Minecraft\"。")
		assert.Contains(t, writer.String(), "水球 對影片 \"Minecraft\" 按讚。")
		writer.Reset()
	})
}

func FakeNewChannel(name string, w io.Writer) *Channel {
	return &Channel{Name: name, Writer: w}
}

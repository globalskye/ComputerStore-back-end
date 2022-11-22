package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type test struct {
	Id         int    `json:"id"`
	Desc       string `json:"desc"`
	Occupation string `json:"occupation"`
}

func (h *Handler) test(c *gin.Context) {

	c.JSON(http.StatusOK, []map[string]interface{}{
		{
			"id":          1,
			"name":        "Ноутбук Lenovo IdeaPad 3 15IGL05 81WQ0059RE",
			"image":       "https://content2.onliner.by/catalog/device/header/557ab393261fb7f79715850e009240da.jpeg",
			"description": "15.6\" 1920 x 1080 IPS, 60 Гц, несенсорный, Intel Celeron N4020 1100 МГц, 4 ГБ DDR4, SSD 256 ГБ, видеокарта встроенная, без ОС, цвет крышки черный",
			"price":       "99,79 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		}, {
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
		{
			"id":          2,
			"name":        "Ноутбук Huawei MateBook D 15 BoB-WAH9Q 53012KRC",
			"image":       "https://content2.onliner.by/catalog/device/header/1bd5893ab33b00756688fb91f9da14e8.jpeg",
			"description": "полноразмерная игровая мышь для ПК, проводная USB, сенсор оптический 16000 dpi, 6 кнопок, колесо с нажатием, цвет черный",
			"price":       "2280,00 p.",
		},
	})
	return

	/*c.JSON(http.StatusOK, "qweJ")*/
}

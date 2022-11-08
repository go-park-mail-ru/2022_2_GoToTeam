package repository

/*
import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/modelsOLD"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"log"
	"sync"
)

var articlesData = []*modelsOLD.Article{
	{
		1,
		"Абракодабра",
		"щршгршгфра фшгрвашгрф пгшр фшгрвзфз ждлж длж  ждьйджлоай ",
		[]string{"tag1", "tag2", "tag3"},
		"Спорт",
		66,
		[]string{"Вася", "Петя", "Иван"},
		"Дядя Семён ехал из города домой. С ним была собака Жучка, Вдруг из леса выскочили волки. Жучка испугалась и прыгнула в сани. У дяди Семёна была хорошая лошадь. Она тоже испугалась и быстро помчалась по дороге. Деревня была близко. Показались огни в окнах. Волки отстали.\n\nУмная лошадь спасла дядю Семена и Жучку.",
	},
	{
		2,
		"МакПакинс",
		"Супер забегаловка",
		[]string{"asd", "qwe"},
		"Food",
		111,
		[]string{"ghj", "kl"},
		"В деревне было много садов. Осенью поспевали яблоки и груши. В садах было много птиц. Они выводили птенцов и целый день кормили их червяками.\n\nРебята разорили гнезда птиц. Птицы улетели из этой деревни. Весной зацвели на яблонях цветы, но червяки забрались в цветы и поели их. Осенью не было на деревьях яблок и груш.\n\nПоняли ребята, что птицы спасали их деревья, но было поздно.\n\nПТИЧИЙ ДВОР\n\nНа ферме большой птичий двор. На дворе гуляют гуси и гусята, утки и утята, куры и цыплята. Птиц кормит птичница бабушка Настя. Ей помогают Таня и Катя. Они кормят гусят, утят и цыплят.",
	},
	{
		3,
		"SevenMeven",
		"Супер забегаловка",
		[]string{"mac", "chik"},
		"Food",
		111,
		[]string{"Mac"},
		"Был у дедушки Степана мёд в горшке. Забрались в горшок муравьи и ели мёд. Дедушка видит, дело плохо. Взял он горшок, привязал веревку и повесил горшок на гвоздь к потолку. А в горшке остался один муравей. Он искал дорогу домой: вылез из горшка на верёвку, потом на потолок. С потолка . га стену, а со стены на пол.\n\nМуравей показал дорогу к горшку другим муравьям. Дедушка Степан снял горшок, а там мёду нет.",
	},
}

type FeedStorage struct {
	articles []*modelsOLD.Article
	mu       sync.RWMutex
	logger   *logger.Logger
}

func NewFeedCustomRepository(logger *logger.Logger) feedComponentInterfaces.FeedRepositoryInterface {
	return &FeedStorage{
		articles: articlesData,
		mu:       sync.RWMutex{},
		logger:   logger,
	}
}

func (o *FeedStorage) PrintArticles() {
	log.Println("Articles in storage:")
	for _, v := range o.articles {
		log.Printf("%#v ", v)
	}
}

func (o *FeedStorage) GetArticles() []*modelsOLD.Article {
	log.Println("Storage GetArticles called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.articles
}


*/

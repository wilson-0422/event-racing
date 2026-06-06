package src

import (
	"event-racing/src/config"
	"event-racing/src/models"
	"event-racing/src/services"
	"log"
)

func Seed() {
	var count int
	config.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if count > 0 {
		log.Println("Seed data already exists, skipping...")
		return
	}

	log.Println("Seeding database...")

	services.CreateUser("admin", "admin123", "admin")
	services.CreateUser("operator", "operator123", "user")

	athletes := []models.Athlete{
		{Name: "张伟", Gender: "男", Age: 25, Team: "北京队", Event: "100米跑", Phone: "13800100001", IDNumber: "110101199901011234"},
		{Name: "李娜", Gender: "女", Age: 23, Team: "上海队", Event: "200米跑", Phone: "13800100002", IDNumber: "310101199802021234"},
		{Name: "王芳", Gender: "女", Age: 27, Team: "广州队", Event: "400米跑", Phone: "13800100003", IDNumber: "440101199703031234"},
		{Name: "刘洋", Gender: "男", Age: 22, Team: "深圳队", Event: "800米跑", Phone: "13800100004", IDNumber: "440301199904041234"},
		{Name: "陈静", Gender: "女", Age: 24, Team: "成都队", Event: "1500米跑", Phone: "13800100005", IDNumber: "510101199805051234"},
		{Name: "杨帆", Gender: "男", Age: 26, Team: "北京队", Event: "跳远", Phone: "13800100006", IDNumber: "110101199806061234"},
		{Name: "赵磊", Gender: "男", Age: 28, Team: "上海队", Event: "跳高", Phone: "13800100007", IDNumber: "310101199607071234"},
		{Name: "黄丽", Gender: "女", Age: 21, Team: "广州队", Event: "铅球", Phone: "13800100008", IDNumber: "440101200108081234"},
		{Name: "周强", Gender: "男", Age: 29, Team: "深圳队", Event: "标枪", Phone: "13800100009", IDNumber: "440301199509091234"},
		{Name: "吴敏", Gender: "女", Age: 23, Team: "成都队", Event: "铁饼", Phone: "13800100010", IDNumber: "510101199910101234"},
		{Name: "孙涛", Gender: "男", Age: 25, Team: "北京队", Event: "100米跑", Phone: "13800100011", IDNumber: "110101199811111234"},
		{Name: "马超", Gender: "男", Age: 24, Team: "上海队", Event: "200米跑", Phone: "13800100012", IDNumber: "310101199912121234"},
		{Name: "朱婷", Gender: "女", Age: 22, Team: "广州队", Event: "跳远", Phone: "13800100013", IDNumber: "440101200001011234"},
		{Name: "胡明", Gender: "男", Age: 27, Team: "深圳队", Event: "800米跑", Phone: "13800100014", IDNumber: "440301199702021234"},
		{Name: "林丹", Gender: "男", Age: 26, Team: "成都队", Event: "400米跑", Phone: "13800100015", IDNumber: "510101199803031234"},
	}
	for i := range athletes {
		services.CreateAthlete(&athletes[i])
	}

	groups := []models.Group{
		{Name: "100米跑预赛A组", Event: "100米跑", Status: "arranged"},
		{Name: "100米跑预赛B组", Event: "100米跑", Status: "arranged"},
		{Name: "200米跑预赛组", Event: "200米跑", Status: "pending"},
		{Name: "跳远决赛组", Event: "跳远", Status: "arranged"},
		{Name: "800米跑决赛组", Event: "800米跑", Status: "pending"},
		{Name: "铅球决赛组", Event: "铅球", Status: "pending"},
	}
	for i := range groups {
		services.CreateGroup(&groups[i])
	}

	services.ArrangeAthletes(1, []int64{1, 11})
	services.ArrangeAthletes(2, []int64{})
	services.ArrangeAthletes(4, []int64{6, 13})

	venues := []models.Venue{
		{Name: "主体育场", Location: "中心区域", Capacity: 5000, Status: "available"},
		{Name: "田径场A区", Location: "东区", Capacity: 2000, Status: "available"},
		{Name: "田径场B区", Location: "西区", Capacity: 2000, Status: "available"},
		{Name: "室内体育馆", Location: "南区", Capacity: 3000, Status: "available"},
		{Name: "投掷场地", Location: "北区", Capacity: 1000, Status: "available"},
	}
	for i := range venues {
		services.CreateVenue(&venues[i])
	}

	scheduleData := []models.Schedule{
		{GroupID: 1, VenueID: 1, StartTime: "2026-06-10 09:00", EndTime: "2026-06-10 10:00", Status: "scheduled"},
		{GroupID: 2, VenueID: 2, StartTime: "2026-06-10 10:00", EndTime: "2026-06-10 11:00", Status: "scheduled"},
		{GroupID: 4, VenueID: 3, StartTime: "2026-06-11 09:00", EndTime: "2026-06-11 11:00", Status: "scheduled"},
		{GroupID: 3, VenueID: 1, StartTime: "2026-06-11 14:00", EndTime: "2026-06-11 15:30", Status: "scheduled"},
		{GroupID: 5, VenueID: 2, StartTime: "2026-06-12 09:00", EndTime: "2026-06-12 10:30", Status: "scheduled"},
	}
	for i := range scheduleData {
		services.CreateSchedule(&scheduleData[i])
	}

	scoreData := []models.Score{
		{AthleteID: 1, ScheduleID: 1, Score: "10.52", Rank: 1, Remark: "个人最佳"},
		{AthleteID: 11, ScheduleID: 1, Score: "10.78", Rank: 2, Remark: ""},
	}
	for i := range scoreData {
		services.CreateScore(&scoreData[i])
	}

	awardData := []models.Award{
		{AthleteID: 1, Event: "100米跑", MedalType: "金牌", CompetitionName: "2026年全国田径锦标赛"},
		{AthleteID: 11, Event: "100米跑", MedalType: "银牌", CompetitionName: "2026年全国田径锦标赛"},
		{AthleteID: 6, Event: "跳远", MedalType: "金牌", CompetitionName: "2026年全国田径锦标赛"},
		{AthleteID: 2, Event: "200米跑", MedalType: "铜牌", CompetitionName: "2025年省运会"},
		{AthleteID: 4, Event: "800米跑", MedalType: "金牌", CompetitionName: "2025年省运会"},
	}
	for i := range awardData {
		services.CreateAward(&awardData[i])
	}

	log.Println("Seed data inserted successfully")
}

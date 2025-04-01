package models

// Response 通用响应结构
type Response struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

// ShuttleList 班车列表响应结构
type ShuttleList struct {
	Count   int64          `json:"count"`
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    []ShuttleRoute `json:"data"`
}

// ShuttleRoute 班车路线信息
type ShuttleRoute struct {
	Pkid               int     `json:"pkid"`
	ID                 string  `json:"id"`
	ShuttleType        int     `json:"shuttle_type"`
	CarNumber          *string `json:"car_number"`
	Name               string  `json:"name"`
	TrainNumber        string  `json:"train_number"`
	ServiceTime        string  `json:"service_time"`
	OriginAddress      string  `json:"origin_address"`
	EndAddress         string  `json:"end_address"`
	IntermediateSite   *string `json:"intermediate_site"`
	OriginTime         string  `json:"origin_time"`
	EndTime            string  `json:"end_time"`
	ReservationNumAble int     `json:"reservation_num_able"`
	Type               int     `json:"type"`
	TeacherTicketPrice string  `json:"teacher_ticket_price"`
	StudentTicketPrice string  `json:"student_ticket_price"`
}

// ShuttleInfo 班车详情响应结构
type ShuttleInfo struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    ShuttleDetail `json:"data"`
}

// ShuttleDetail 班车详细信息
type ShuttleDetail struct {
	Pkid               int     `json:"pkid"`
	ID                 int64   `json:"id"`
	WeekDay            string  `json:"week_day"`
	Name               string  `json:"name"`
	Type               int     `json:"type"`
	ShuttleType        int     `json:"shuttle_type"`
	TrainNumber        string  `json:"train_number"`
	Status             int     `json:"status"`
	CarNumber          *string `json:"car_number"`
	MapImage           string  `json:"map_image"`
	OrgUid             int     `json:"org_uid"`
	OrgName            *string `json:"org_name"`
	DeptID             int     `json:"dept_id"`
	ServiceTime        string  `json:"service_time"`
	OriginTime         string  `json:"origin_time"`
	EndTime            string  `json:"end_time"`
	OriginLongitude    int     `json:"origin_longitude"`
	OriginLatitude     int     `json:"origin_latitude"`
	OriginAddress      string  `json:"origin_address"`
	EndLongitude       int     `json:"end_longitude"`
	EndLatitude        int     `json:"end_latitude"`
	EndAddress         string  `json:"end_address"`
	IntermediateSite   *string `json:"intermediate_site"`
	IsFree             int     `json:"is_free"`
	TeacherTicketPrice string  `json:"teacher_ticket_price"`
	StudentTicketPrice string  `json:"student_ticket_price"`
	IsDeleted          int     `json:"is_deleted"`
	BatchID            int     `json:"batch_id"`
	ReservationNumAble int     `json:"reservation_num_able"`
	VehicleStatus      int     `json:"vehicle_status"`
	GmtModified        string  `json:"gmt_modified"`
	GmtCreated         string  `json:"gmt_created"`
}

// ReservedSeats 座位预定状态响应结构
type ReservedSeats struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    ReservationState `json:"data"`
}

// ReservationState 预定状态信息
type ReservationState struct {
	ReservedCount      int      `json:"reserved_count"`
	ReservationNum     int      `json:"reservation_num"`
	ReservedSeatNumber []string `json:"reserved_seat_number"`
	IsFull             int      `json:"is_full"`
}

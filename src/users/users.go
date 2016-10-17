package users

// CREATE TABLE students ("period","section_id","grd","enter_date","student_name","ethnic_code","perm_id","birth_date","home_language","phone","nm_state_id","trk","course_id","course_title","staff_name","room_name","student_school_year_gu","student_gu","gen","female","track","male","audit_class","meeting_days" text);

type (
	Student struct {
		ID          string
		StudentName string
		Grade       string
	}
	
)

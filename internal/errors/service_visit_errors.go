package errors

// Service Visit service errors
var (
    ErrInvalidServiceVisitCreateRequest = NewWithCode("INVALID_SERVICE_VISIT_CREATE", "invalid service visit create request", "درخواست ساخت بازدید سرویس معتبر نیست")
    ErrInvalidServiceVisitUpdateRequest = NewWithCode("INVALID_SERVICE_VISIT_UPDATE", "invalid service visit update request", "درخواست به روز رسانی بازدید سرویس معتبر نیست")
    ErrInvalidServiceVisitID            = NewWithCode("INVALID_SERVICE_VISIT_ID", "invalid service visit id", "شناسه بازدید سرویس نامعتبر است")
    ErrFailedToCreateServiceVisit       = NewWithCode("CREATE_SERVICE_VISIT_FAILED", "failed to create service visit", "خطای ساخت بازدید سرویس")
    ErrFailedToGetServiceVisit          = NewWithCode("GET_SERVICE_VISIT_FAILED", "failed to get service visit", "خطای دریافت بازدید سرویس")
    ErrFailedToListServiceVisits        = NewWithCode("LIST_SERVICE_VISITS_FAILED", "failed to list service visits", "خطای فهرست بازدیدهای سرویس")
    ErrFailedToUpdateServiceVisit       = NewWithCode("UPDATE_SERVICE_VISIT_FAILED", "failed to update service visit", "خطای به روز رسانی بازدید سرویس")
    ErrFailedToDeleteServiceVisit       = NewWithCode("DELETE_SERVICE_VISIT_FAILED", "failed to delete service visit", "خطای حذف بازدید سرویس")
) 
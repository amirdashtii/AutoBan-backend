package errors

// Oil Change service errors
var (
    ErrInvalidOilChangeCreateRequest = NewWithCode("INVALID_OIL_CHANGE_CREATE", "invalid oil change create request", "درخواست ساخت تعویض روغن معتبر نیست")
    ErrInvalidOilChangeUpdateRequest = NewWithCode("INVALID_OIL_CHANGE_UPDATE", "invalid oil change update request", "درخواست به روز رسانی تعویض روغن معتبر نیست")
    ErrInvalidOilChangeID            = NewWithCode("INVALID_OIL_CHANGE_ID", "invalid oil change id", "شناسه تعویض روغن نامعتبر است")
    ErrFailedToCreateOilChange       = NewWithCode("CREATE_OIL_CHANGE_FAILED", "failed to create oil change", "خطای ساخت تعویض روغن")
    ErrFailedToGetOilChange          = NewWithCode("GET_OIL_CHANGE_FAILED", "failed to get oil change", "خطای دریافت تعویض روغن")
    ErrFailedToListOilChanges        = NewWithCode("LIST_OIL_CHANGES_FAILED", "failed to list oil changes", "خطای فهرست تعویض روغن")
    ErrFailedToUpdateOilChange       = NewWithCode("UPDATE_OIL_CHANGE_FAILED", "failed to update oil change", "خطای به روز رسانی تعویض روغن")
    ErrFailedToDeleteOilChange       = NewWithCode("DELETE_OIL_CHANGE_FAILED", "failed to delete oil change", "خطای حذف تعویض روغن")
    ErrInvalidDate                   = NewWithCode("INVALID_DATE", "invalid date format", "فرمت تاریخ نامعتبر است")
    ErrUserVehicleIDRequired         = NewWithCode("USER_VEHICLE_ID_REQUIRED", "user vehicle id is required", "شناسه وسیله نقلیه کاربر الزامی است")
    ErrOilChangeNotOwned             = NewWithCode("OIL_CHANGE_NOT_OWNED", "oil change not owned", "تعویض روغن متعلق به کاربر نیست")
) 
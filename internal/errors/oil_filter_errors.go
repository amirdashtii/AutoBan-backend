package errors

// Oil Filter service errors
var (
    ErrInvalidOilFilterCreateRequest = NewWithCode("INVALID_OIL_FILTER_CREATE", "invalid oil filter create request", "درخواست ساخت تعویض فیلتر روغن معتبر نیست")
    ErrInvalidOilFilterUpdateRequest = NewWithCode("INVALID_OIL_FILTER_UPDATE", "invalid oil filter update request", "درخواست به روز رسانی تعویض فیلتر روغن معتبر نیست")
    ErrInvalidOilFilterID            = NewWithCode("INVALID_OIL_FILTER_ID", "invalid oil filter id", "شناسه تعویض فیلتر روغن نامعتبر است")
    ErrFailedToCreateOilFilter       = NewWithCode("CREATE_OIL_FILTER_FAILED", "failed to create oil filter", "خطای ساخت تعویض فیلتر روغن")
    ErrFailedToGetOilFilter          = NewWithCode("GET_OIL_FILTER_FAILED", "failed to get oil filter", "خطای دریافت تعویض فیلتر روغن")
    ErrOilFilterNotOwned             = NewWithCode("OIL_FILTER_NOT_OWNED", "oil filter not owned", "تعویض فیلتر روغن متعلق به کاربر نیست")
    ErrFailedToListOilFilters        = NewWithCode("LIST_OIL_FILTERS_FAILED", "failed to list oil filters", "خطای فهرست تعویض فیلتر روغن")
    ErrFailedToUpdateOilFilter       = NewWithCode("UPDATE_OIL_FILTER_FAILED", "failed to update oil filter", "خطای به روز رسانی تعویض فیلتر روغن")
    ErrFailedToDeleteOilFilter       = NewWithCode("DELETE_OIL_FILTER_FAILED", "failed to delete oil filter", "خطای حذف تعویض فیلتر روغن")
) 
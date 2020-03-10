package v8runnner

////////////////////////////////////////////////////////
// Доступные опции
//
//func WithStartParams(params string) types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("/C", params)
//	}
//}
//
//func WithUnlockCode(uc string) types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("/UC", uc)
//	}
//}
//
//func WithUpdateDBCfg() types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("/UpdateDBCfg", true)
//	}
//}
//
//func WithUpdateDBCfgOptions(options *UpdateDBCfgOptions) types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("/UpdateDBCfg", options)
//	}
//}
//
//func WithExtension(ext string) types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("-Extension", ext)
//	}
//}
//
//func WithManagedApplication() types.UserOption {
//	return func(o types.Optioned) {
//		o.SetOption("/RunModeManagedApplication", true)
//	}
//}

//func WithCredentials(user, password string) types.UserOption {
//	return func(o types.Optioned) {
//
//		if len(user) == 0 {
//			return
//		}
//
//		o.SetOption("/U", user)
//
//		if len(password) > 0 {
//			o.SetOption("/P", user)
//		}
//
//	}
//}

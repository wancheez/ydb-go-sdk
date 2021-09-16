// Code generated by gtrace. DO NOT EDIT.

package table

import (
	"context"
	"time"
)

// Compose returns a new ClientTrace which has functional fields composed
// both from t and x.
func (t ClientTrace) Compose(x ClientTrace) (ret ClientTrace) {
	switch {
	case t.OnCreateSession == nil:
		ret.OnCreateSession = x.OnCreateSession
	case x.OnCreateSession == nil:
		ret.OnCreateSession = t.OnCreateSession
	default:
		h1 := t.OnCreateSession
		h2 := x.OnCreateSession
		ret.OnCreateSession = func(c CreateSessionStartInfo) func(CreateSessionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CreateSessionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnKeepAlive == nil:
		ret.OnKeepAlive = x.OnKeepAlive
	case x.OnKeepAlive == nil:
		ret.OnKeepAlive = t.OnKeepAlive
	default:
		h1 := t.OnKeepAlive
		h2 := x.OnKeepAlive
		ret.OnKeepAlive = func(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
			r1 := h1(k)
			r2 := h2(k)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(k KeepAliveDoneInfo) {
					r1(k)
					r2(k)
				}
			}
		}
	}
	switch {
	case t.OnDeleteSession == nil:
		ret.OnDeleteSession = x.OnDeleteSession
	case x.OnDeleteSession == nil:
		ret.OnDeleteSession = t.OnDeleteSession
	default:
		h1 := t.OnDeleteSession
		h2 := x.OnDeleteSession
		ret.OnDeleteSession = func(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DeleteSessionDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	switch {
	case t.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = x.OnPrepareDataQuery
	case x.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = t.OnPrepareDataQuery
	default:
		h1 := t.OnPrepareDataQuery
		h2 := x.OnPrepareDataQuery
		ret.OnPrepareDataQuery = func(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PrepareDataQueryDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = x.OnExecuteDataQuery
	case x.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = t.OnExecuteDataQuery
	default:
		h1 := t.OnExecuteDataQuery
		h2 := x.OnExecuteDataQuery
		ret.OnExecuteDataQuery = func(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
			r1 := h1(e)
			r2 := h2(e)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(e ExecuteDataQueryDoneInfo) {
					r1(e)
					r2(e)
				}
			}
		}
	}
	switch {
	case t.OnStreamReadTable == nil:
		ret.OnStreamReadTable = x.OnStreamReadTable
	case x.OnStreamReadTable == nil:
		ret.OnStreamReadTable = t.OnStreamReadTable
	default:
		h1 := t.OnStreamReadTable
		h2 := x.OnStreamReadTable
		ret.OnStreamReadTable = func(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamReadTableDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = x.OnStreamExecuteScanQuery
	case x.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = t.OnStreamExecuteScanQuery
	default:
		h1 := t.OnStreamExecuteScanQuery
		h2 := x.OnStreamExecuteScanQuery
		ret.OnStreamExecuteScanQuery = func(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamExecuteScanQueryDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnBeginTransaction == nil:
		ret.OnBeginTransaction = x.OnBeginTransaction
	case x.OnBeginTransaction == nil:
		ret.OnBeginTransaction = t.OnBeginTransaction
	default:
		h1 := t.OnBeginTransaction
		h2 := x.OnBeginTransaction
		ret.OnBeginTransaction = func(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
			r1 := h1(b)
			r2 := h2(b)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(b BeginTransactionDoneInfo) {
					r1(b)
					r2(b)
				}
			}
		}
	}
	switch {
	case t.OnCommitTransaction == nil:
		ret.OnCommitTransaction = x.OnCommitTransaction
	case x.OnCommitTransaction == nil:
		ret.OnCommitTransaction = t.OnCommitTransaction
	default:
		h1 := t.OnCommitTransaction
		h2 := x.OnCommitTransaction
		ret.OnCommitTransaction = func(c CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CommitTransactionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = x.OnRollbackTransaction
	case x.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = t.OnRollbackTransaction
	default:
		h1 := t.OnRollbackTransaction
		h2 := x.OnRollbackTransaction
		ret.OnRollbackTransaction = func(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RollbackTransactionDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	return ret
}
func (t ClientTrace) onCreateSession(c1 CreateSessionStartInfo) func(CreateSessionDoneInfo) {
	fn := t.OnCreateSession
	if fn == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onKeepAlive(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
	fn := t.OnKeepAlive
	if fn == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	res := fn(k)
	if res == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onDeleteSession(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
	fn := t.OnDeleteSession
	if fn == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onPrepareDataQuery(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
	fn := t.OnPrepareDataQuery
	if fn == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onExecuteDataQuery(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
	fn := t.OnExecuteDataQuery
	if fn == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	res := fn(e)
	if res == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onStreamReadTable(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
	fn := t.OnStreamReadTable
	if fn == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onStreamExecuteScanQuery(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
	fn := t.OnStreamExecuteScanQuery
	if fn == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onBeginTransaction(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
	fn := t.OnBeginTransaction
	if fn == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	res := fn(b)
	if res == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onCommitTransaction(c1 CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
	fn := t.OnCommitTransaction
	if fn == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onRollbackTransaction(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
	fn := t.OnRollbackTransaction
	if fn == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	return res
}

// Compose returns a new SessionPoolTrace which has functional fields composed
// both from t and x.
func (t SessionPoolTrace) Compose(x SessionPoolTrace) (ret SessionPoolTrace) {
	switch {
	case t.OnCreate == nil:
		ret.OnCreate = x.OnCreate
	case x.OnCreate == nil:
		ret.OnCreate = t.OnCreate
	default:
		h1 := t.OnCreate
		h2 := x.OnCreate
		ret.OnCreate = func(s SessionPoolCreateStartInfo) func(SessionPoolCreateDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolCreateDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnGet == nil:
		ret.OnGet = x.OnGet
	case x.OnGet == nil:
		ret.OnGet = t.OnGet
	default:
		h1 := t.OnGet
		h2 := x.OnGet
		ret.OnGet = func(s SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolGetDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnWait == nil:
		ret.OnWait = x.OnWait
	case x.OnWait == nil:
		ret.OnWait = t.OnWait
	default:
		h1 := t.OnWait
		h2 := x.OnWait
		ret.OnWait = func(s SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolWaitDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnTake == nil:
		ret.OnTake = x.OnTake
	case x.OnTake == nil:
		ret.OnTake = t.OnTake
	default:
		h1 := t.OnTake
		h2 := x.OnTake
		ret.OnTake = func(s SessionPoolTakeStartInfo) func(SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
					r11 := r1(s)
					r21 := r2(s)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(s SessionPoolTakeDoneInfo) {
							r11(s)
							r21(s)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnPut == nil:
		ret.OnPut = x.OnPut
	case x.OnPut == nil:
		ret.OnPut = t.OnPut
	default:
		h1 := t.OnPut
		h2 := x.OnPut
		ret.OnPut = func(s SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolPutDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnCloseSession == nil:
		ret.OnCloseSession = x.OnCloseSession
	case x.OnCloseSession == nil:
		ret.OnCloseSession = t.OnCloseSession
	default:
		h1 := t.OnCloseSession
		h2 := x.OnCloseSession
		ret.OnCloseSession = func(s SessionPoolCloseSessionStartInfo) func(SessionPoolCloseSessionDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolCloseSessionDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnClose == nil:
		ret.OnClose = x.OnClose
	case x.OnClose == nil:
		ret.OnClose = t.OnClose
	default:
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(s SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolCloseDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	return ret
}
func (t SessionPoolTrace) onCreate(s SessionPoolCreateStartInfo) func(SessionPoolCreateDoneInfo) {
	fn := t.OnCreate
	if fn == nil {
		return func(SessionPoolCreateDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolCreateDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onGet(s SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo) {
	fn := t.OnGet
	if fn == nil {
		return func(SessionPoolGetDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onWait(s SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo) {
	fn := t.OnWait
	if fn == nil {
		return func(SessionPoolWaitDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolWaitDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onTake(s SessionPoolTakeStartInfo) func(SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
	fn := t.OnTake
	if fn == nil {
		return func(SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
			return func(SessionPoolTakeDoneInfo) {
				return
			}
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
			return func(SessionPoolTakeDoneInfo) {
				return
			}
		}
	}
	return func(s SessionPoolTakeWaitInfo) func(SessionPoolTakeDoneInfo) {
		res := res(s)
		if res == nil {
			return func(SessionPoolTakeDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t SessionPoolTrace) onPut(s SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo) {
	fn := t.OnPut
	if fn == nil {
		return func(SessionPoolPutDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onCloseSession(s SessionPoolCloseSessionStartInfo) func(SessionPoolCloseSessionDoneInfo) {
	fn := t.OnCloseSession
	if fn == nil {
		return func(SessionPoolCloseSessionDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolCloseSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onClose(s SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(SessionPoolCloseDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolCloseDoneInfo) {
			return
		}
	}
	return res
}

// Compose returns a new createSessionTrace which has functional fields composed
// both from t and x.
func (t createSessionTrace) Compose(x createSessionTrace) (ret createSessionTrace) {
	switch {
	case t.OnCheckEnoughSpace == nil:
		ret.OnCheckEnoughSpace = x.OnCheckEnoughSpace
	case x.OnCheckEnoughSpace == nil:
		ret.OnCheckEnoughSpace = t.OnCheckEnoughSpace
	default:
		h1 := t.OnCheckEnoughSpace
		h2 := x.OnCheckEnoughSpace
		ret.OnCheckEnoughSpace = func(enoughSpace bool) {
			h1(enoughSpace)
			h2(enoughSpace)
		}
	}
	switch {
	case t.OnCreateSessionGoroutineStart == nil:
		ret.OnCreateSessionGoroutineStart = x.OnCreateSessionGoroutineStart
	case x.OnCreateSessionGoroutineStart == nil:
		ret.OnCreateSessionGoroutineStart = t.OnCreateSessionGoroutineStart
	default:
		h1 := t.OnCreateSessionGoroutineStart
		h2 := x.OnCreateSessionGoroutineStart
		ret.OnCreateSessionGoroutineStart = func() func(createSessionResult) {
			r1 := h1()
			r2 := h2()
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r createSessionResult) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnStartSelect == nil:
		ret.OnStartSelect = x.OnStartSelect
	case x.OnStartSelect == nil:
		ret.OnStartSelect = t.OnStartSelect
	default:
		h1 := t.OnStartSelect
		h2 := x.OnStartSelect
		ret.OnStartSelect = func() {
			h1()
			h2()
		}
	}
	switch {
	case t.OnReadResult == nil:
		ret.OnReadResult = x.OnReadResult
	case x.OnReadResult == nil:
		ret.OnReadResult = t.OnReadResult
	default:
		h1 := t.OnReadResult
		h2 := x.OnReadResult
		ret.OnReadResult = func(r createSessionResult) {
			h1(r)
			h2(r)
		}
	}
	switch {
	case t.OnContextDone == nil:
		ret.OnContextDone = x.OnContextDone
	case x.OnContextDone == nil:
		ret.OnContextDone = t.OnContextDone
	default:
		h1 := t.OnContextDone
		h2 := x.OnContextDone
		ret.OnContextDone = func() {
			h1()
			h2()
		}
	}
	switch {
	case t.OnPutSession == nil:
		ret.OnPutSession = x.OnPutSession
	case x.OnPutSession == nil:
		ret.OnPutSession = t.OnPutSession
	default:
		h1 := t.OnPutSession
		h2 := x.OnPutSession
		ret.OnPutSession = func(session *Session, err error) {
			h1(session, err)
			h2(session, err)
		}
	}
	return ret
}
func (t createSessionTrace) onCheckEnoughSpace(enoughSpace bool) {
	fn := t.OnCheckEnoughSpace
	if fn == nil {
		return
	}
	fn(enoughSpace)
}
func (t createSessionTrace) onCreateSessionGoroutineStart() func(createSessionResult) {
	fn := t.OnCreateSessionGoroutineStart
	if fn == nil {
		return func(createSessionResult) {
			return
		}
	}
	res := fn()
	if res == nil {
		return func(createSessionResult) {
			return
		}
	}
	return res
}
func (t createSessionTrace) onStartSelect() {
	fn := t.OnStartSelect
	if fn == nil {
		return
	}
	fn()
}
func (t createSessionTrace) onReadResult(r createSessionResult) {
	fn := t.OnReadResult
	if fn == nil {
		return
	}
	fn(r)
}
func (t createSessionTrace) onContextDone() {
	fn := t.OnContextDone
	if fn == nil {
		return
	}
	fn()
}
func (t createSessionTrace) onPutSession(session *Session, err error) {
	fn := t.OnPutSession
	if fn == nil {
		return
	}
	fn(session, err)
}
func clientTraceOnCreateSession(t ClientTrace, c context.Context) func(_ context.Context, _ *Session, endpoint string, latency time.Duration, _ error) {
	var p CreateSessionStartInfo
	p.Context = c
	res := t.onCreateSession(p)
	return func(c context.Context, s *Session, endpoint string, latency time.Duration, e error) {
		var p CreateSessionDoneInfo
		p.Context = c
		p.Session = s
		p.Endpoint = endpoint
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func clientTraceOnKeepAlive(t ClientTrace, c context.Context, s *Session) func(context.Context, *Session, SessionInfo, error) {
	var p KeepAliveStartInfo
	p.Context = c
	p.Session = s
	res := t.onKeepAlive(p)
	return func(c context.Context, s *Session, s1 SessionInfo, e error) {
		var p KeepAliveDoneInfo
		p.Context = c
		p.Session = s
		p.SessionInfo = s1
		p.Error = e
		res(p)
	}
}
func clientTraceOnDeleteSession(t ClientTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, latency time.Duration, _ error) {
	var p DeleteSessionStartInfo
	p.Context = c
	p.Session = s
	res := t.onDeleteSession(p)
	return func(c context.Context, s *Session, latency time.Duration, e error) {
		var p DeleteSessionDoneInfo
		p.Context = c
		p.Session = s
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func clientTraceOnPrepareDataQuery(t ClientTrace, c context.Context, s *Session, query string) func(_ context.Context, _ *Session, query string, result *DataQuery, cached bool, _ error) {
	var p PrepareDataQueryStartInfo
	p.Context = c
	p.Session = s
	p.Query = query
	res := t.onPrepareDataQuery(p)
	return func(c context.Context, s *Session, query string, result *DataQuery, cached bool, e error) {
		var p PrepareDataQueryDoneInfo
		p.Context = c
		p.Session = s
		p.Query = query
		p.Result = result
		p.Cached = cached
		p.Error = e
		res(p)
	}
}
func clientTraceOnExecuteDataQuery(t ClientTrace, c context.Context, s *Session, txID string, query *DataQuery, parameters *QueryParameters) func(_ context.Context, _ *Session, txID string, query *DataQuery, parameters *QueryParameters, prepared bool, _ *Result, _ error) {
	var p ExecuteDataQueryStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	p.Query = query
	p.Parameters = parameters
	res := t.onExecuteDataQuery(p)
	return func(c context.Context, s *Session, txID string, query *DataQuery, parameters *QueryParameters, prepared bool, r *Result, e error) {
		var p ExecuteDataQueryDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Query = query
		p.Parameters = parameters
		p.Prepared = prepared
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnStreamReadTable(t ClientTrace, c context.Context, s *Session) func(context.Context, *Session, *Result, error) {
	var p StreamReadTableStartInfo
	p.Context = c
	p.Session = s
	res := t.onStreamReadTable(p)
	return func(c context.Context, s *Session, r *Result, e error) {
		var p StreamReadTableDoneInfo
		p.Context = c
		p.Session = s
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnStreamExecuteScanQuery(t ClientTrace, c context.Context, s *Session, query *DataQuery, parameters *QueryParameters) func(_ context.Context, _ *Session, query *DataQuery, parameters *QueryParameters, _ *Result, _ error) {
	var p StreamExecuteScanQueryStartInfo
	p.Context = c
	p.Session = s
	p.Query = query
	p.Parameters = parameters
	res := t.onStreamExecuteScanQuery(p)
	return func(c context.Context, s *Session, query *DataQuery, parameters *QueryParameters, r *Result, e error) {
		var p StreamExecuteScanQueryDoneInfo
		p.Context = c
		p.Session = s
		p.Query = query
		p.Parameters = parameters
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnBeginTransaction(t ClientTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, txID string, _ error) {
	var p BeginTransactionStartInfo
	p.Context = c
	p.Session = s
	res := t.onBeginTransaction(p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p BeginTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func clientTraceOnCommitTransaction(t ClientTrace, c context.Context, s *Session, txID string) func(_ context.Context, _ *Session, txID string, _ error) {
	var p CommitTransactionStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	res := t.onCommitTransaction(p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p CommitTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func clientTraceOnRollbackTransaction(t ClientTrace, c context.Context, s *Session, txID string) func(_ context.Context, _ *Session, txID string, _ error) {
	var p RollbackTransactionStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	res := t.onRollbackTransaction(p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p RollbackTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnCreate(t SessionPoolTrace, c context.Context) func(context.Context, *Session, error) {
	var p SessionPoolCreateStartInfo
	p.Context = c
	res := t.onCreate(p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolCreateDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnGet(t SessionPoolTrace, c context.Context) func(_ context.Context, _ *Session, latency time.Duration, retryAttempts int, _ error) {
	var p SessionPoolGetStartInfo
	p.Context = c
	res := t.onGet(p)
	return func(c context.Context, s *Session, latency time.Duration, retryAttempts int, e error) {
		var p SessionPoolGetDoneInfo
		p.Context = c
		p.Session = s
		p.Latency = latency
		p.RetryAttempts = retryAttempts
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnWait(t SessionPoolTrace, c context.Context) func(context.Context, *Session, error) {
	var p SessionPoolWaitStartInfo
	p.Context = c
	res := t.onWait(p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolWaitDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnTake(t SessionPoolTrace, c context.Context, s *Session) func(context.Context, *Session) func(_ context.Context, _ *Session, took bool, _ error) {
	var p SessionPoolTakeStartInfo
	p.Context = c
	p.Session = s
	res := t.onTake(p)
	return func(c context.Context, s *Session) func(context.Context, *Session, bool, error) {
		var p SessionPoolTakeWaitInfo
		p.Context = c
		p.Session = s
		res := res(p)
		return func(c context.Context, s *Session, took bool, e error) {
			var p SessionPoolTakeDoneInfo
			p.Context = c
			p.Session = s
			p.Took = took
			p.Error = e
			res(p)
		}
	}
}
func sessionPoolTraceOnPut(t SessionPoolTrace, c context.Context, s *Session) func(context.Context, *Session, error) {
	var p SessionPoolPutStartInfo
	p.Context = c
	p.Session = s
	res := t.onPut(p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolPutDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnCloseSession(t SessionPoolTrace, c context.Context, s *Session) func(context.Context, *Session, error) {
	var p SessionPoolCloseSessionStartInfo
	p.Context = c
	p.Session = s
	res := t.onCloseSession(p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolCloseSessionDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnClose(t SessionPoolTrace, c context.Context) func(context.Context, error) {
	var p SessionPoolCloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(c context.Context, e error) {
		var p SessionPoolCloseDoneInfo
		p.Context = c
		p.Error = e
		res(p)
	}
}
func createSessionTraceOnCheckEnoughSpace(t createSessionTrace, enoughSpace bool) {
	t.onCheckEnoughSpace(enoughSpace)
}
func createSessionTraceOnCreateSessionGoroutineStart(t createSessionTrace) func() {
	res := t.onCreateSessionGoroutineStart()
	return func() {
		var p createSessionResult
		res(p)
	}
}
func createSessionTraceOnStartSelect(t createSessionTrace) {
	t.onStartSelect()
}
func createSessionTraceOnReadResult(t createSessionTrace) {
	var p createSessionResult
	t.onReadResult(p)
}
func createSessionTraceOnContextDone(t createSessionTrace) {
	t.onContextDone()
}
func createSessionTraceOnPutSession(t createSessionTrace, session *Session, err error) {
	t.onPutSession(session, err)
}
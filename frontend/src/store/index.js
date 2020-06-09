import { createStore } from 'redux'
import rootReducer from './reducers'

const initialStore = {}

export default createStore(rootReducer, initialStore)

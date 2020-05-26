import React, { useReducer, useEffect } from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import { Paper, makeStyles } from '@material-ui/core'
import TopBar from '../../components/appBar/topBar'
import PersonGridList from '../../molecules/person/personGridList'

const useStyles = makeStyles({
  root: {
    paddingTop: '16px',
  }
})

const App = () => {
  const classes = useStyles()
  const initialState = {
    localData: null,
    singerData: null,
    peersData: null,
    error: null,
    loaded: false,
    fetching: false,
  }
  const reducer = (state, newState) => ({ ...state, ...newState })
  const [state, setState] = useReducer(reducer, initialState)

  async function fetchData() {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true, video: true })
    const localData = {
      id: stream.id,
      stream
    }
    const peersData = []
    const singerData = {}

    // error?
    if (!localData.stream) {
      return setState({
        localData,
        singerData,
        peersData,
        error: true,
        loaded: true,
        fetching: false
      })
    }

    //no error
    setState({
      localData,
      singerData,
      peersData,
      error: null,
      loaded: true,
      fetching: false,
    })
  }

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <Paper className={classes.root}>
          {state.loaded ? <PersonGridList singer={state.singerData} local={state.localData} peers={state.peersData} /> : null }
        </Paper>
      </main>
    </>
  )
}

export default App

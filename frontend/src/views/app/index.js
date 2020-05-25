import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import GroupGridList from '../../molecules/group/groupGridList'

export default function App() {
  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        <GroupGridList />
      </main>
    </>
  )
}

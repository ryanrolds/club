import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'

export default function App() {
  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        {/* Paper or Main Page Component Here */}
      </main>
    </>
  )
}

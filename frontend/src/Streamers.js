import React from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
    streamerContext: {
        flex: 1,
    }
})

export default function Streamers() {
    const classes = useStyles()
    return (
        <React.Fragment>
           <div className={classes}>hi</div>
        </React.Fragment>
    )
}
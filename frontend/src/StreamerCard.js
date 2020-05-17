import React from 'react'
import { makeStyles, Card, CardMedia } from '@material-ui/core'

export default class StreamerCard extends React.Component {
    constructor(props) {
        super(props)
        this.streamRef = React.createRef()
    }

    useStyles() {
        return makeStyles(() => ({
            card: {
                height: '100%',
                display: 'flex',
                flexDirection: 'column',
            },
            cardMedia: {
                paddingTop: '100%', // 16:9
            },
        }))
    }

    render() {
        const classes = this.useStyles();

        return (
            <Card className={classes.card}>
                <CardMedia
                    component="video"
                    className={classes.cardMedia}
                    ref={this.streamRef}
                    title="A streamer"
                />
            </Card>
        )
    }

    componentDidMount() {
        this.updateVideoStream()
    }

    componentDidUpdate() {
        this.updateVideoStream()
    }

    updateVideoStream() {
        if (this.streamRef.current.srcObject !== this.props.stream) {
            this.streamRef.current.srcObject = this.props.stream
            this.streamRef.current.autoplay = true
            this.streamRef.current.muted = this.props.muted
        }
    }
}
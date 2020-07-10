import React from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button'

const useStyles = makeStyles({
  gridItem: (colSize) => ({
    'grid-column': `${colSize}fr`,
    'text-align': 'center',
  }),
})

const GroupGridItem = ({ name, onClick }) => {
  const classes = useStyles(2)

  return (
    <div className={classes.gridItem}>
      <Button onClick={onClick}>
        Join&nbsp;
        {name}
      </Button>
    </div>
  )
}

GroupGridItem.propTypes = {
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
}

export default GroupGridItem

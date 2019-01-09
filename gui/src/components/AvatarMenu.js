import React from 'react'
import { withStyles } from '@material-ui/core/styles'
import { Link } from 'react-static'
import { useHooks, useState, useRef } from 'use-react-hooks'
import Popper from '@material-ui/core/Popper'
import Grow from '@material-ui/core/Grow'
import Fab from '@material-ui/core/Fab'
import Paper from '@material-ui/core/Paper'
import MenuItem from '@material-ui/core/MenuItem'
import MenuList from '@material-ui/core/MenuList'
import ClickAwayListener from '@material-ui/core/ClickAwayListener'
import Avatar from '@material-ui/core/Avatar'
import Typography from '@material-ui/core/Typography'
import { grey } from '@material-ui/core/colors'
import face from 'images/face.svg'

const styles = theme => {
  return {
    avatar: {
      width: 21,
      height: 25
    },
    menuItem: {
      '&:hover': {
        backgroundColor: 'transparent'
      }
    },
    link: {
      color: grey[600],
      textDecoration: 'none',
      '&:hover': {
        color: grey[900]
      }
    }
  }
}

const AvatarMenu = useHooks(({ classes }) => {
  const anchorEl = useRef(null)
  const [open, setOpenState] = useState(false)
  const handleToggle = () => setOpenState(!open)

  const handleClose = event => {
    if (anchorEl.current.contains(event.target)) {
      return
    }

    setOpenState(false)
  }

  return (
    <React.Fragment>
      <Fab
        size='small'
        color='primary'
        aria-label='Profile'
        className={classes.button}
        buttonRef={anchorEl}
        aria-owns={open ? 'menu-list-grow' : undefined}
        aria-haspopup='true'
        onClick={handleToggle}
      >
        <Avatar alt='Profile' src={face} className={classes.avatar} />
      </Fab>
      <Popper open={open} anchorEl={anchorEl.current} transition disablePortal>
        {({ TransitionProps, placement }) => (
          <Grow
            {...TransitionProps}
            id='menu-list-grow'
            style={{ transformOrigin: placement === 'bottom' ? 'center top' : 'center bottom' }}
          >
            <Paper>
              <ClickAwayListener onClickAway={handleClose}>
                <MenuList>
                  <MenuItem onClick={handleClose} className={classes.menuItem}>
                    <Link to='/signout' className={classes.link}>
                      <Typography
                        variant='body1'
                        className={classes.link}
                      >
                        Log out
                      </Typography>
                    </Link>
                  </MenuItem>
                </MenuList>
              </ClickAwayListener>
            </Paper>
          </Grow>
        )}
      </Popper>
    </React.Fragment>
  )
})

export default withStyles(styles)(AvatarMenu)

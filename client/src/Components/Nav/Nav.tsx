import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import Link from '@mui/material/Link';
import History from '@mui/icons-material/History';
import Bookmarks from '@mui/icons-material/Bookmarks';
import VideoGameAsset from '@mui/icons-material/VideogameAsset';

export default function Nav() {
  return (
      <AppBar
        position="fixed"
        style={{ background: 'transparent', boxShadow: 'none' }}
        // sx={{ borderBottom: 1, borderColor: 'divider' }}
      >
        <Toolbar>
          {/* <Themes /> */}

          <Typography sx={{ flexGrow: 1 }} />

          <Link href="/game">
            <IconButton
              size="large"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
            >
              <VideoGameAsset sx={{ fontSize: 30, color: "slateblue" }} />
            </IconButton>
          </Link>

          <Link href="/word">
            <IconButton
              size="large"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
            >
              <Bookmarks sx={{ fontSize: 30, color: "cadetblue" }} />
            </IconButton>
          </Link>

          <Link href="/act">
            <IconButton
              size="large"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
            >
              <History sx={{ fontSize: 30, color: "darkorange" }} />
            </IconButton>
          </Link>
          
          {/* <Button color="inherit"><VideoGameAsset style={{ color: "slateblue" }} /></Button> */}
        </Toolbar>
      </AppBar>

  );
}

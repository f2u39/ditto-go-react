import { useEffect, useState } from "react";
import dayjs, { Dayjs } from 'dayjs';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { AppBar, Box, Divider, Grid, IconButton, Link, Stack, Toolbar, Tooltip, Typography } from '@mui/material';

import PostAddIcon from '@mui/icons-material/PostAdd';
import TimerIcon from '@mui/icons-material/Timer';
import DateRangeIcon from '@mui/icons-material/DateRange';

import { Git, Controller } from 'react-bootstrap-icons';

import TodayIcon from '@mui/icons-material/Today';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import SportsEsportsIcon from '@mui/icons-material/SportsEsports';
import GitHubIcon from '@mui/icons-material/GitHub';

import FormatListNumberedRtlIcon from '@mui/icons-material/FormatListNumberedRtl';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import TitleIcon from '@mui/icons-material/Title';

import ArrowCircleLeftIcon from '@mui/icons-material/ArrowCircleLeft';
import ArrowCircleRightIcon from '@mui/icons-material/ArrowCircleRight';

import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from "@mui/x-date-pickers";

import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';

export default function Act() {
    const [period, setPeriod] = useState("Daily")
    const handlePeriodChange = (event: React.SyntheticEvent, newPeriod: string) => {
        setPeriod(newPeriod)
    }

    const [date, setDate] = useState<Dayjs | null>(
        // dayjs(new Date()).format('YYYYMMDD'),
        dayjs(new Date())
    );
    const [tempDate, setTempDate] = useState<Dayjs | null>(
        dayjs(new Date()),
    );
    const [openCalendar, setOpenCalendar] = useState(false);
    const handleCalendarOpen = () => { setOpenCalendar(true) };
    const handleCalendarClose = () => { setOpenCalendar(false) };
    const handleChangeTempDate = (newValue: Dayjs | null) => {
        setTempDate(newValue);
    };

    const handleUpdateDate = () => {
        setDate(tempDate);
        setOpenCalendar(false);
    };

    const [acts, setActs] = useState({
        details: [],
        summary: [],
    });

    useEffect(() => {
        fetch(`/act?date=${date}&period=${period}`)
            .then(resp => resp.json())
            .then(data => {
                if (data != null) {
                    setActs(data)
                }
            })
    }, [date, period]);

    const details = Array.isArray(acts.details) ? acts.details : [];
    const summary: any = acts.summary ? acts.summary : [];

    return (
        <Box sx={{ width: '100%' }}>
            <TabContext value={period}>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <TabList indicatorColor="secondary" onChange={handlePeriodChange} centered>
                        <Tab
                            icon={
                                <TodayIcon fontSize="large" />
                            }
                            value="Daily"
                        />

                        <Tab
                            icon={
                                <CalendarMonthIcon fontSize="large" />
                            }
                            value="Monthly"
                        />
                    </TabList>
                </Box>

                <TabPanel value={period}>
                    <Grid
                        container
                        display="flex"
                        justifyContent="center"
                        alignItems="center"
                        sx={{ pt: 5 }}
                        xs={12}
                    >
                        <Grid item xs={8}>
                            <Box sx={{ flexGrow: 1 }}>
                                <AppBar position="static" style={{ background: 'transparent', boxShadow: 'none' }}>
                                    <Toolbar>
                                        <Typography sx={{ flexGrow: 1 }} />

                                        <IconButton
                                            size="large"
                                            aria-controls="menu-appbar"
                                            aria-haspopup="true"
                                            color="inherit"
                                        >
                                            <PostAddIcon sx={{ fontSize: 30, color: "white" }} />
                                        </IconButton>

                                        <IconButton
                                            size="large"
                                            aria-controls="menu-appbar"
                                            aria-haspopup="true"
                                            color="inherit"
                                        >
                                            <TimerIcon sx={{ fontSize: 30, color: "white" }} />
                                        </IconButton>

                                        <IconButton
                                            size="large"
                                            aria-controls="menu-appbar"
                                            aria-haspopup="true"
                                            color="inherit"
                                            onClick={handleCalendarOpen}
                                        >
                                            <DateRangeIcon sx={{ fontSize: 30, color: "white" }} />
                                        </IconButton>
                                        <Dialog open={openCalendar} onClose={handleCalendarClose}>
                                            <DialogTitle>Select a date</DialogTitle>
                                            <DialogContent>
                                                <LocalizationProvider dateAdapter={AdapterDayjs}>
                                                    <DesktopDatePicker
                                                        // autoFocus
                                                        inputFormat="MM/DD/YYYY"
                                                        value={tempDate}
                                                        onChange={handleChangeTempDate}
                                                        renderInput={(params) => <TextField {...params} />}
                                                    />
                                                </LocalizationProvider>
                                            </DialogContent>
                                            <DialogActions>
                                                <Button onClick={handleCalendarClose}>Cancel</Button>
                                                <Button onClick={handleUpdateDate}>Search</Button>
                                            </DialogActions>
                                        </Dialog>

                                        <Typography sx={{ flexGrow: 1 }} />
                                    </Toolbar>
                                </AppBar>
                            </Box>
                            <TableContainer sx={{ borderRadius: 1, border: 2 }}>
                                <Toolbar>
                                    <IconButton>
                                        <ArrowCircleLeftIcon />
                                    </IconButton>

                                    {/* <Typography
                                        align="center"
                                        sx={{ flex: '1 1 100%' }}
                                        variant="h6"
                                        component="div"
                                        onClick={handleCalendarOpen}
                                    >
                                        {dayjs(date).format('DD MMM (ddd) YYYY')}
                                        <DateRangeIcon fontSize="large" />
                                    </Typography> */}

                                    {/* <Typography
                                        align="center"
                                        sx={{ flex: '1 1 100%' }}
                                        variant="h6"
                                        component="div"
                                    > */}
                                    {/* <Stack
                                            direction="row"
                                            alignItems="center"
                                        >
                                            <Typography>
                                                {dayjs(date).format('DD MMM (ddd) YYYY')}
                                            </Typography>
                                            <IconButton>
                                                <DateRangeIcon fontSize="large" onClick={handleCalendarOpen} />
                                            </IconButton>
                                        </Stack> */}

                                    <Typography
                                        align="center"
                                        sx={{ flex: '1 1 100%' }}
                                        variant="h6"
                                        component="div"
                                    >
                                        <Grid container xs={12} direction="row" alignItems="center">
                                            <DateRangeIcon /> example
                                        </Grid>

                                    </Typography>

                                    {/* </Typography> */}


                                    <Tooltip title="Date">
                                        <IconButton>
                                            <ArrowCircleRightIcon />
                                        </IconButton>
                                    </Tooltip>
                                </Toolbar>
                                <Table>
                                    <TableHead>
                                        <TableRow>
                                            <TableCell align="center" style={{ width: 50 }}><FormatListNumberedRtlIcon /></TableCell>
                                            <TableCell align="center" style={{ width: 80 }}><AccessTimeIcon /></TableCell>
                                            <TableCell align="left"><TitleIcon /></TableCell>
                                            <TableCell style={{ width: 120 }}></TableCell>
                                        </TableRow>
                                    </TableHead>

                                    <TableBody>
                                        {(details).map(
                                            (detail: any) => {
                                                return (
                                                    <TableRow
                                                        key={detail.id}
                                                        sx={{ '&:last-child td, &:last-child th': { border: 0, fontSize: 15 } }}
                                                    >
                                                        {detail.act.type === 'Gaming' ?
                                                            <TableCell align="center"><Typography color="lightpink"><SportsEsportsIcon /></Typography></TableCell> :
                                                            <TableCell align="center"><Typography color="mediumpurple"><GitHubIcon /></Typography></TableCell>}

                                                        {detail.act.type === 'Gaming' ?
                                                            <TableCell align="center"><Typography color="lightpink">{detail.act.duration}</Typography></TableCell> :
                                                            <TableCell align="center"><Typography color="mediumpurple">{detail.act.duration}</Typography></TableCell>}

                                                        {detail.act.type === 'Gaming' ?
                                                            <TableCell colSpan={2} align="left"><Typography color="lightpink">{detail.game[0].title}</Typography></TableCell> :
                                                            <TableCell colSpan={2}></TableCell>}
                                                    </TableRow>
                                                )
                                            }
                                        )}
                                        <TableRow>
                                            <TableCell colSpan={2}></TableCell>
                                            <TableCell align="right"><Typography color="mediumpurple"><Controller size={23} /></Typography></TableCell>
                                            <TableCell align="right"><Typography color="mediumpurple">{summary.pgm_hour} h {summary.pgm_min} m</Typography></TableCell>
                                        </TableRow>
                                        <TableRow>
                                            <TableCell colSpan={2}></TableCell>
                                            <TableCell align="right"><Typography color="lightpink"><Git size={23} /></Typography></TableCell>
                                            <TableCell align="right"><Typography color="lightpink">{summary.game_hour} h {summary.game_min} m</Typography></TableCell>
                                        </TableRow>
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        </Grid>
                    </Grid>
                </TabPanel>
            </TabContext>
        </Box>
    )
}


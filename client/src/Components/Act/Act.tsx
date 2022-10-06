import { useEffect, useState } from "react";
import dayjs, { Dayjs } from 'dayjs';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { AppBar, Box, Divider, Grid, IconButton, Link, Toolbar, Typography } from '@mui/material';

import PostAddIcon from '@mui/icons-material/PostAdd';
import TimerIcon from '@mui/icons-material/Timer';
import DateRangeIcon from '@mui/icons-material/DateRange';

import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import ReorderIcon from '@mui/icons-material/Reorder';

import TodayIcon from '@mui/icons-material/Today';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';

import SportsEsportsIcon from '@mui/icons-material/SportsEsports';
import GitHubIcon from '@mui/icons-material/GitHub';

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
        dayjs(new Date()),
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
        // fetchActs();
        fetch(`/act?period=${period}`)
            .then(resp => resp.json())
            .then(data => {
                // console.log(data)
                if (data != null) {
                    setActs(data)
                }
            })
    }, [period]);

    // const dailyActs = Array.isArray(acts.daily_acts) ? acts.daily_acts : [];
    // const DailyTableRows = (dailyActs).map(
    //     (detail: any) => {
    //         return (
    //             <TableRow key={detail.id} sx={{ '&:last-child td, &:last-child th': { border: 0, fontSize: 15 } }}>
    //                 {
    //                     detail.act.type === 'Gaming' ?
    //                         <TableCell align="center"><SportsEsportsIcon /></TableCell> :
    //                         <TableCell align="center"><GitHubIcon /></TableCell>
    //                 }
    //                 <TableCell align="center">{detail.act.duration}</TableCell>
    //                 {
    //                     detail.game.length === 1 ?
    //                         <TableCell align="left">{detail.game[0].title}</TableCell> :
    //                         <TableCell></TableCell>
    //                 }
    //             </TableRow>
    //         )
    //     }
    // );

    // const daySum: any = acts.day_sum != null ? acts.day_sum : [];
    // const DaySumTableRows =
    //     <>
    //         <TableRow>
    //             <TableCell align="right" colSpan={2}><SportsEsportsIcon sx={{ color: "cadetblue" }} /></TableCell>
    //             {daySum.length === 0 ?
    //                 <TableCell></TableCell> :
    //                 <TableCell align="left">{daySum.game_hour}h {daySum.game_min}m</TableCell>
    //             }
    //         </TableRow>
    //         <TableRow>
    //             <TableCell align="right" colSpan={2}><GitHubIcon sx={{ color: "cadetblue" }} /></TableCell>
    //             {daySum.length === 0 ?
    //                 <TableCell></TableCell> :
    //                 <TableCell align="left">{daySum.pgm_hour}h {daySum.pgm_min}m</TableCell>
    //             }
    //         </TableRow>
    //     </>;

    // const monthlyActs = Array.isArray(acts.monthly_acts) ? acts.monthly_acts : [];
    // const MonthlyTableRows = (monthlyActs).map(
    //     (detail: any) => {
    //         return (
    //             <TableRow key={detail.id} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
    //                 {
    //                     detail.act.type === 'Gaming' ?
    //                         <TableCell align="left"><SportsEsportsIcon /></TableCell> :
    //                         <TableCell><GitHubIcon /></TableCell>
    //                 }
    //                 <TableCell align="center">{detail.act.duration}</TableCell>
    //                 {
    //                     detail.game.length === 1 ?
    //                         <TableCell align="left">{detail.game[0].title}</TableCell> :
    //                         <TableCell></TableCell>
    //                 }
    //             </TableRow>
    //         )
    //     }
    // );
    // const monthSum: any = acts.month_sum != null ? acts.month_sum : [];
    // const MonthSumTableRows =
    //     <>
    //         <TableRow>
    //             <TableCell align="right" colSpan={2}><SportsEsportsIcon sx={{ color: "cadetblue" }} /></TableCell>
    //             {monthSum.length === 0 ?
    //                 <TableCell></TableCell> :
    //                 <TableCell align="left">{monthSum.game_hour}h {monthSum.game_min}m</TableCell>
    //             }
    //         </TableRow>
    //         <TableRow>
    //             <TableCell align="right" colSpan={2}><GitHubIcon sx={{ color: "cadetblue" }} /></TableCell>
    //             {monthSum.length === 0 ?
    //                 <TableCell></TableCell> :
    //                 <TableCell align="left">{monthSum.pgm_hour}h {monthSum.pgm_min}m</TableCell>
    //             }
    //         </TableRow>
    //     </>;

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
                                            <PostAddIcon sx={{ fontSize: 30, color: "#F8C8DC" }} />
                                        </IconButton>

                                        <IconButton
                                            size="large"
                                            aria-controls="menu-appbar"
                                            aria-haspopup="true"
                                            color="inherit"
                                        >
                                            <TimerIcon sx={{ fontSize: 30, color: "#F8C8DC" }} />
                                        </IconButton>

                                        <IconButton
                                            size="large"
                                            aria-controls="menu-appbar"
                                            aria-haspopup="true"
                                            color="inherit"
                                            onClick={handleCalendarOpen}
                                        >
                                            <DateRangeIcon sx={{ fontSize: 30, color: "#F8C8DC" }} />
                                        </IconButton>
                                        <Dialog open={openCalendar} onClose={handleCalendarClose}>
                                            <DialogTitle>Select a date</DialogTitle>
                                            <DialogContent>
                                                {/* <DialogContentText></DialogContentText> */}
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
                                <Table>
                                    <TableHead>
                                        <TableRow>
                                            <TableCell colSpan={3}>
                                                <div style={{ display: 'flex', alignItems: 'center', flexWrap: 'wrap' }}>
                                                    <TodayIcon />
                                                    <span>Daily</span>
                                                </div>
                                            </TableCell>
                                        </TableRow>
                                    </TableHead>
                                </Table>
                            </TableContainer>
                        </Grid>
                    </Grid>
                </TabPanel>
            </TabContext>
        </Box>
    )
}


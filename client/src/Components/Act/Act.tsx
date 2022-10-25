import { useCallback, useEffect, useReducer, useState } from "react";
import dayjs, { Dayjs } from 'dayjs';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import { AppBar, Box, FormControl, FormControlLabel, FormGroup, FormLabel, Grid, IconButton, InputLabel, Link, MenuItem, Select, Stack, Switch, Toolbar, Tooltip, Typography } from '@mui/material';
import PostAddIcon from '@mui/icons-material/PostAdd';
import TimerIcon from '@mui/icons-material/Timer';
import DateRangeIcon from '@mui/icons-material/DateRange';
import { Git as GitIcon } from 'react-bootstrap-icons';
import { Controller as ControllerIcon } from 'react-bootstrap-icons';
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
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import Tab from '@mui/material/Tab';
import { useForm, Controller } from "react-hook-form";

export default function Act() {
    const [date, setDate] = useState<Dayjs | null>(dayjs(new Date()));
    const handleUpdateDate = (newValue: Dayjs | null) => {
        setDate(newValue)
        setOpenCalendar(false)
    }
    const handlePreviousDate = () => {
        setDate(dayjs(date).add(-1, 'day'))
    }
    const handleNextDate = () => {
        setDate(dayjs(date).add(1, 'day'))
    }

    const [tempDate, setTempDate] = useState<Dayjs | null>(dayjs(new Date()));
    const handleChangeTempDate = (newValue: Dayjs | null) => {
        setTempDate(newValue);
    }

    const [openNewActivity, setOpenNewActivity] = useState(false)
    const handleNewActivityOpen = () => { setOpenNewActivity(true) }
    const handleNewActivityClose = () => {
        setOpenNewActivity(false)
        setFormValues(defaultValues)
    }

    const [openCalendar, setOpenCalendar] = useState(false)
    const handleCalendarOpen = () => { setOpenCalendar(true) }
    const handleCalendarClose = () => { setOpenCalendar(false) }

    const [acts, setActs] = useState({
        day_details: [],
        day_summary: [],
        month_details: [],
        month_summary: [],
        playing_games: [],
    })

    useEffect(() => {
        fetchActs()
    }, [date])

    function fetchActs() {
        fetch(`/act?date=${dayjs(date).format('YYYYMMDD')}`)
            .then(resp => resp.json())
            .then(data => {
                if (data != null) {
                    setActs(data)
                }
            })
    }

    const dayDetails = Array.isArray(acts.day_details) ? acts.day_details : []
    const daySummary: any = acts.day_summary ? acts.day_summary : []
    const monDetails = Array.isArray(acts.month_details) ? acts.month_details : []
    const monSummary: any = acts.month_summary ? acts.month_summary : []
    const playingGames = Array.isArray(acts.playing_games) ? acts.playing_games : []

    const defaultValues = {
        type: 'Gaming',
        date: dayjs(new Date()).format('YYYYMMDD'),
        duration: '',
        gameId: '',
    }

    const [formValues, setFormValues] = useState(defaultValues)

    const handleInputChange = (e: { target: { name: any; value: any; } }) => {
        const { name, value } = e.target;
        setFormValues({
            ...formValues,
            [name]: value,
        })
    }

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        // console.log(formValues)

        fetch("/act/create", {
            method: "POST",
            body: JSON.stringify(formValues),
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then(response => response.json())
            // .then(response => console.log("Success:", JSON.stringify(response)))
            .then(() => {
                handleNewActivityClose()
                fetchActs()
            })
            .catch(error => console.error("Error:", error))
    }

    return (
        <Box sx={{ width: '100%' }}>
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
                                    <PostAddIcon onClick={handleNewActivityOpen} sx={{ fontSize: 35, color: "#0461B1" }} />
                                </IconButton>

                                <IconButton
                                    size="large"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    color="inherit"
                                >
                                    <TimerIcon sx={{ fontSize: 35, color: "#0461B1" }} />
                                </IconButton>

                                <Typography sx={{ flexGrow: 1 }} />
                            </Toolbar>
                        </AppBar>
                    </Box>
                    <TableContainer sx={{ border: 1, borderRadius: 1, borderColor: 'gray' }}>
                        <Toolbar sx={{ borderBottom: 1, borderColor: 'gray' }}>
                            <Tooltip title="Previous date">
                                <IconButton onClick={handlePreviousDate}>
                                    <ArrowCircleLeftIcon />
                                </IconButton>
                            </Tooltip>

                            <Grid
                                container
                                spacing={0}
                                direction="row"
                                alignItems="center"
                                justifyContent="center"
                            >
                                <Grid item>
                                    <Grid container direction="row" alignItems="center">
                                        {dayjs(date).format('DD MMM (ddd) YYYY')}
                                        <Tooltip title="Pick date">
                                            <IconButton onClick={handleCalendarOpen}><DateRangeIcon fontSize="large" /></IconButton>
                                        </Tooltip>
                                    </Grid>
                                </Grid>
                            </Grid>

                            <Tooltip title="Next date">
                                <IconButton onClick={handleNextDate}>
                                    <ArrowCircleRightIcon />
                                </IconButton>
                            </Tooltip>
                        </Toolbar>

                        <Table size="small">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="center" style={{ width: 50 }}><FormatListNumberedRtlIcon /></TableCell>
                                    <TableCell align="center" style={{ width: 80 }}><AccessTimeIcon /></TableCell>
                                    <TableCell align="left"><TitleIcon /></TableCell>
                                    <TableCell style={{ width: 120 }}></TableCell>
                                </TableRow>
                            </TableHead>

                            <TableBody>
                                <TableRow sx={{ borderTop: 1, borderColor: 'gray' }}>
                                    <TableCell colSpan={4} align="center">📆Daily</TableCell>
                                </TableRow>

                                {(dayDetails).map(
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
                                    <TableCell align="right"><Typography color="lightpink"><ControllerIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right"><Typography color="lightpink">{daySummary.game_hour} h {daySummary.game_min} m</Typography></TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right"><Typography color="mediumpurple"><GitIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right"><Typography color="mediumpurple">{daySummary.pgm_hour} h {daySummary.pgm_min} m</Typography></TableCell>
                                </TableRow>

                                <TableRow sx={{ borderTop: 1, borderColor: 'gray' }}>
                                    <TableCell colSpan={4} align="center">📅Monthly</TableCell>
                                </TableRow>

                                {(monDetails).map(
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
                                    <TableCell align="right"><Typography color="lightpink"><ControllerIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right"><Typography color="lightpink">{monSummary.game_hour} h {monSummary.game_min} m</Typography></TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell colSpan={2}></TableCell>
                                    <TableCell align="right"><Typography color="mediumpurple"><GitIcon size={23} /></Typography></TableCell>
                                    <TableCell align="right"><Typography color="mediumpurple">{monSummary.pgm_hour} h {monSummary.pgm_min} m</Typography></TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Grid>
            </Grid>

            <Dialog open={openCalendar} onClose={handleCalendarClose}>
                <DialogTitle>Select a date</DialogTitle>
                <DialogContent>
                    <LocalizationProvider dateAdapter={AdapterDayjs}>
                        <DesktopDatePicker
                            inputFormat={"MM/DD/YYYY"}
                            value={tempDate}
                            onChange={handleUpdateDate}
                            renderInput={(params) => <TextField {...params} />}
                        />
                    </LocalizationProvider>
                </DialogContent>
            </Dialog>

            <Dialog
                open={openNewActivity}
                onClose={handleNewActivityClose}
            >
                <DialogTitle align="center">New Activity</DialogTitle>
                <DialogContent>
                    <form onSubmit={handleSubmit}>
                        <FormControl sx={{ mt: 2, minWidth: 500 }}>
                            <InputLabel htmlFor="type">Type</InputLabel>
                            <Select
                                name="type"
                                label="Type"
                                value={formValues.type}
                                onChange={handleInputChange}
                            >
                                <MenuItem value="Gaming">Gaming</MenuItem>
                                <MenuItem value="Programming">Programming</MenuItem>
                            </Select>
                        </FormControl>

                        <FormControl sx={{ mt: 2, minWidth: 500 }}>
                            <LocalizationProvider dateAdapter={AdapterDayjs}>
                                <DatePicker
                                    label="Date"
                                    inputFormat={"MM/DD/YYYY"}
                                    value={tempDate}
                                    onChange={handleUpdateDate}
                                    renderInput={(params) =>
                                        <TextField {...params}
                                            name="date"
                                            value={formValues.date}
                                            onChange={handleInputChange}
                                        />
                                    }
                                />
                            </LocalizationProvider>
                        </FormControl>

                        <FormControl sx={{ mt: 2, minWidth: 500 }}>
                            {/* <InputLabel htmlFor="duration">Duration</InputLabel> */}
                            <TextField
                                name="duration"
                                label="Duration"
                                type="number"
                                value={formValues.duration}
                                onChange={handleInputChange}
                                InputProps={{
                                    inputProps: { min: 0 }
                                }}
                            />
                        </FormControl>

                        <FormControl sx={{ mt: 2, minWidth: 500 }}>
                            <InputLabel htmlFor="type">Game</InputLabel>
                            <Select
                                name="gameId"
                                label="Game"
                                value={formValues.gameId}
                                inputProps={{
                                    name: 'gameId',
                                }}
                                onChange={handleInputChange}
                            >
                                {playingGames.map((game: any, index) => {
                                    return (
                                        <MenuItem key={index} value={game.id}>{game.title}</MenuItem>
                                    )
                                })}
                            </Select>
                        </FormControl>

                        <FormControl sx={{ mt: 2 }}>
                            <Stack direction="row" spacing={2} justifyContent="flex-end">
                                <Button onClick={handleNewActivityClose}>Cancel</Button>
                                <Button type="submit">Submit</Button>
                            </Stack>
                        </FormControl>
                    </form>
                </DialogContent>
            </Dialog>
        </Box>
    )
}

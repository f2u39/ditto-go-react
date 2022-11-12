import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import TuneIcon from '@mui/icons-material/Tune';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';
import { Badge, Box, Grid, InputAdornment, Tabs, TextField, Tooltip } from '@mui/material';
import Pagination from '@mui/material/Pagination';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import { CheckSquareFill, Tablet, PcDisplay, NintendoSwitch, Playstation, Xbox } from 'react-bootstrap-icons';
import { Code, CodeSlash } from 'react-bootstrap-icons';
import { Battery, BatteryCharging, BatteryFull } from 'react-bootstrap-icons';
import { useEffect, useState } from 'react';

interface ExpandMoreProps extends IconButtonProps {
    expand: boolean;
}

const ExpandMore = styled((props: ExpandMoreProps) => {
    const { expand, ...other } = props;
    return <IconButton {...other} />;
})(({ theme, expand }) => ({
    transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
        duration: theme.transitions.duration.shortest,
    }),
}));

export default function Game() {
    const [details, setDetails] = useState<Detail[]>([]);
    const [platform, setPlatform] = useState('All');
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [status, setStatus] = useState("Playing")
    const [playedCount, setPlayedCount] = useState(0);
    const [playingCount, setPlayingCount] = useState(0);
    const [blockingCount, setBlockingCount] = useState(0);

    useEffect(() => {
        fetch(`/api/game/status/${status}/${platform}/${page}`)
            .then(resp => resp.json())
            .then(data => {
                // console.log(data["details"])
                if (data["details"] != null) {
                    setDetails(data["details"])
                } else {
                    setDetails([])
                }

                setTotalPages(data["total_pages"])
            })

        fetch(`/api/game/counts`)
            .then(resp => resp.json())
            .then(data => {
                setPlayedCount(data["played_cnt"])
                setPlayingCount(data["playing_cnt"])
                setBlockingCount(data["blocking_cnt"])
            })
    }, [status, platform, page]);

    const handlePageChange = (event: React.ChangeEvent<unknown>, value: number) => {
        setPage(value);
    };

    const handleStatusChange = (event: React.SyntheticEvent, newStatus: string) => {
        setPage(1)
        setStatus(newStatus)
    };

    const handlePlatformChange = (event: React.SyntheticEvent, newValue: string) => {
        setPage(1)
        setPlatform(newValue)
    };

    const handleStartGame = (id: string) => {
        fetch(`/api/act/watch/start?id=${id}`)
    }

    return (
        <Box sx={{ width: '100%' }}>
            <TabContext value={status}>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <TabList indicatorColor="secondary" onChange={handleStatusChange} centered>
                        <Tab
                            icon={
                                <Badge badgeContent={playedCount} color="primary">
                                    <BatteryFull fontSize="30" color="white" />
                                </Badge>
                            }
                            value="Played"
                        />
                        <Tab
                            icon={
                                <Badge badgeContent={playingCount} color="success">
                                    <BatteryCharging fontSize="30" color="green" />
                                </Badge>
                            }
                            value="Playing"
                        />
                        <Tab
                            icon={
                                <Badge badgeContent={blockingCount} color="error">
                                    <Battery fontSize="30" color="red" />
                                </Badge>
                            }
                            value="Blocking"
                        />
                    </TabList>
                </Box>

                <TabPanel value={status}>
                    <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        sx={{ display: 'inline-flex' }}
                    >
                        <Grid
                            item
                            sx={{ m: -3, borderRight: 1, borderColor: 'divider' }}
                        >
                            <Tabs
                                variant="fullWidth"
                                orientation="vertical"
                                value={platform}
                                onChange={handlePlatformChange}
                            >
                                <Tab sx={{ mt: 6 }} icon={<CheckSquareFill size={30} />} value="All" />
                                <Tab sx={{ mt: 2 }} icon={<PcDisplay color="orange" size={30} />} value="PC" />
                                <Tab sx={{ mt: 2 }} icon={<Playstation color="#2E6DB4" size={30} />} value="PlayStation" />
                                <Tab sx={{ mt: 2 }} icon={<NintendoSwitch color="#E60012" size={30} />} value="Nintendo Switch" />
                                <Tab sx={{ mt: 2 }} icon={<Xbox color="#107C10" size={30} />} value="Xbox" />
                                <Tab sx={{ mt: 2 }} icon={<Tablet color="#730073" size={30} />} value="Mobile" />
                            </Tabs>
                        </Grid>

                        <Grid item xs={10}>
                            <Grid
                                container
                            >
                                {details.map((element, i) => (
                                    <Card
                                        sx={{ ml: 3, mt: 3, maxWidth: 250 }}
                                        key={element.game.id}
                                    >
                                        <CardMedia
                                            component="img"
                                            height="250"
                                            image={"/assets/images/games/" + element.game.id + ".webp"}
                                        />
                                        <CardContent>
                                            <Typography variant="subtitle1" align="center" color="text.secondary">
                                                {element.game.title}
                                            </Typography>
                                        </CardContent>
                                        <CardActions sx={{ mt: -1 }} disableSpacing>
                                            <Tooltip title="Property">
                                                <IconButton>
                                                    <TuneIcon />
                                                </IconButton>
                                            </Tooltip>

                                            <Tooltip title="Start">
                                                <IconButton onClick={e => handleStartGame(element.game.id)}>
                                                    <PlayCircleOutlineIcon />
                                                </IconButton>
                                            </Tooltip>
                                        </CardActions>
                                            <CardContent sx={{ mt: -4 }}>
                                                <Box sx={{
                                                    mx: "auto",
                                                    display: 'flex',
                                                    flexDirection: 'column',
                                                    alignItems: 'center',
                                                    '& > *': {
                                                        m: 1,
                                                    },
                                                }}
                                                >
                                                    <TextField
                                                        fullWidth
                                                        size="small"
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { textAlign: 'right' },
                                                            readOnly: true,
                                                        }}
                                                        value={element.game.rating}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    {element.game.platform === 'Mobile' ? <Tablet /> : <></>}
                                                                    {element.game.platform === 'PC' ? <PcDisplay /> : <></>}
                                                                    {element.game.platform === 'Playstation' ? <Playstation /> : <></>}
                                                                    {element.game.platform === 'Nintendo Switch' ? <NintendoSwitch /> : <></>}
                                                                    {element.game.platform === 'Xbox' ? <Xbox /> : <></>}
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />
                                                    <TextField
                                                        fullWidth
                                                        size="small"
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { textAlign: 'right' },
                                                            readOnly: true,
                                                        }}
                                                        value={element.developer.name}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    <Code />
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />

                                                    <TextField
                                                        fullWidth
                                                        size="small"
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { textAlign: 'right' },
                                                            readOnly: true,
                                                        }}
                                                        value={element.publisher.name}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    <CodeSlash />
                                                                </InputAdornment>
                                                            )
                                                        }}
                                                    />

                                                    <TextField
                                                        fullWidth
                                                        size="small"
                                                        sx={{ pt: 1 }}
                                                        inputProps={{
                                                            style: { textAlign: 'right' },
                                                            readOnly: true,
                                                        }}
                                                        value={element.play_hour}
                                                        InputProps={{
                                                            startAdornment: (
                                                                <InputAdornment position="start">
                                                                    {element.game.status === 'Played' ? <BatteryFull /> : <></>}
                                                                    {element.game.status === 'Playing' ? <BatteryCharging /> : <></>}
                                                                    {element.game.status === 'Blocking' ? <Battery /> : <></>}
                                                                </InputAdornment>
                                                            ),
                                                            endAdornment: (
                                                                <InputAdornment position="end">Hour(s)</InputAdornment>
                                                            )
                                                        }}
                                                    />
                                                </Box>
                                            </CardContent>
                                    </Card>
                                ))
                                }
                            </Grid>
                        </Grid>
                        <Grid xs={12} sx={{ pt: 3, pb: 3 }}>
                            <Box
                                display="flex"
                                justifyContent="center"
                                alignItems="center"
                            >
                                <Pagination
                                    count={totalPages}
                                    page={page}
                                    onChange={handlePageChange}
                                    variant="outlined"
                                    color="secondary" />
                            </Box>
                        </Grid>
                    </Grid>
                </TabPanel>
            </TabContext>
        </Box>
    );
}


interface Detail {
    game: Game,
    developer: Developer,
    publisher: Publisher,
    play_hour: 0,
    play_min: 0,
}

interface Game {
    id: string,
    title: string,
    genre: string,
    platform: string,
    developer_id: string,
    publisher_id: string,
    status: string,
    playtime: string,
    rating: string,
}

interface Developer {
    id: string,
    name: string,
}

interface Publisher {
    id: string,
    name: string,
}
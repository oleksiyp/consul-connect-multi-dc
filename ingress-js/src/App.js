import React from 'react';
import './App.scss';
import 'react-circular-progressbar/dist/styles.css';
import ReactSpeedometer from "react-d3-speedometer";
import {CircularProgressbar} from "react-circular-progressbar";

class App extends React.Component {
    n = 800;
    minRate = 500;


    constructor(props) {
        super(props);
        this.state = {
            rate: 0,
            dcTags: {},
            services: {},
            oks: {},
            count: 0
        };
    }

    componentDidMount() {
        this.issueRequest();
    }

    issueRequest = () => {
        function itemAt(hash, key, el) {
            if (!(key in hash)) {
                hash[key] = el;
            }
            return hash[key]
        }

        const startTime = new Date().getTime();
        this.lastRequest = startTime;
        fetch("/traffic/svc1?n=" + this.n)
            .then(it => it.ok ? it.json() : Promise.reject("non OK response"))
            .then(json => {
                const endTime = new Date().getTime();
                let count = 0;
                const services = {};
                const dcTags = {};
                for (const name of Object.keys(json.oks)) {
                    count += json.oks[name];
                    const [dc, serviceName, tag] = name.split("/");
                    itemAt(dcTags, dc, {})[tag] = true;
                    itemAt(services, serviceName, true);
                }
                const rate = 1000 * count / (endTime - startTime);
                this.setState(state => {
                    const tt = this.merge(state.dcTags, dcTags);
                    const svcs = this.merge(state.services, services);

                    return ({rate, dcTags: tt, services: svcs, oks: json.oks, count})
                });
            }, () => this.setState({rate: 0, oks: {}, count: 0}))
            .finally(this.scheduleNewRequest);
    };

    merge(first, second) {
        const result = {...first};
        for (const it of Object.keys(second)) {
            if (!(it in result)) {
                result[it] = second[it];
            }
        }
        return result;
    }

    scheduleNewRequest = () => {
        const currentTime = new Date().getTime();
        const delta = currentTime - this.lastRequest;
        if (delta < this.minRate) {
            window.setTimeout(this.issueRequest, this.minRate - delta);
        } else {
            this.issueRequest();
        }
    };

    render() {
        function sortSet(map) {
            const keys = [...Object.keys(map)];
            keys.sort();
            return keys;
        }

        return <div className="row">
            <div>
                <ReactSpeedometer
                    width={200}
                    height={150}
                    ringWidth={30}
                    value={this.state.rate}
                    startColor="#33CC33"
                    endColor="#FF471A"
                    segments={6}
                    currentValueText="RPS: ${value}"
                    needleTransitionDuration={3000}
                    maxValue={1200}
                    valueFormat=".1f"
                />
            </div>
            <div>
                <table>
                    <thead>
                    <tr>
                        <th rowSpan={2}>Service</th>
                        {sortSet(this.state.dcTags).map(dc =>
                            <th key={dc} colSpan={sortSet(this.state.dcTags[dc]).length}>{dc}</th>)}
                    </tr>
                    <tr>
                        {sortSet(this.state.dcTags).flatMap(dc =>
                            sortSet(this.state.dcTags[dc]).map(tag =>
                                <th key={dc + "/" + tag}>{tag}</th>))
                        }
                    </tr>
                    </thead>
                    <tbody>
                    {sortSet(this.state.services).map(service =>
                        <tr key={service}>
                            <td>{service}</td>
                            {sortSet(this.state.dcTags).flatMap(dc =>
                                sortSet(this.state.dcTags[dc]).map(tag =>
                                    <td key={dc + "/" + tag}>{this.renderGauge(service, dc, tag)}</td>))
                            }
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
        </div>
    }

    renderGauge(service, dc, tag) {
        const amount = this.state.oks[dc + "/" + service + "/" + tag] || 0;
        const percentage = this.state.count < 0.001 ? 0 : amount / this.state.count * 100.0;
        return <CircularProgressbar
            className="gauge"
            value={percentage}
            text={percentage.toFixed(1) + "%"}
            styles={{
                path: {stroke: "green", transitionDuration: "0.5s"},
                text: {fill: "green", fontSize: "14px"}
            }}
        />
    }
}

export default App;

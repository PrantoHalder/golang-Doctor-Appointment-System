<h3>Doctor Appointment System</h3>
<h4>Features</h4>
<ol>
<li>Admin can create and manage doctor types</li>
<li>Admin can manage users (doctor,patient)</li>
<li>Guest users can register as patients</li>
<li>The user (doctor,patient)can login</li>
<li>the doctor can manage appoinments</li>
<li>the doctor can manage his/her schedule</li>
<li>the patient ca search for doctors and filer doctor by doctor types and book an appointment</li>
<li>patient can see his or /her appointment</li>
</ol>
<h4>Installation</h4>
<h5>Usermgm Config</h5>

<p><b>[Server]</b></p>
<p>host = 127.0.0.1</p>
<p>port = "4000"</p>
<p><b>[database]</b></p>
<p>host = "127.0.0.1"</p>
<p>port = 5432</p>
<p>dbname = postgres</p>
<p>password = 123</p>
<p>sslmode = disabled</p>
<h5>CMS Config</h5>
<p><b>[server]</b></p>
<p>host = "127.0.0.1"</p>
<p>port = "3000"</p>
<p><b>[session]</b></p>
<p>lifetime = 24</p>
<p>idletime = 20</p>
<p><b>[database]</b></p>
<p>host = 127.0.0.1</p>
<p>Port = 5432</p>
<p>dbname = doctor_cms</p>
<p>user = postgress</p>
<p>password = 123</p>
<p>sslmode = disables</p>
<p><b>[usermgm]</b></p>
<p>url = "127.0.0.1:4000"</p>

<h4>Requirements</h4>
<ul>
<li>Go should be istalled</li>
<li>Postgress should be installed</li>
<li>Browser should be istalled to serve localhost</li>
<li>A Edittor should be installed to run the project .Ex: Visual studio</li>
</ul>
<h4>Tables</h4>
<h5>Users table</h5>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>First name </td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Last name</td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Username</td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>unique</td>
            </tr>
            <tr>
                <td>Email</td>
                <td>(varchar(50))</td>
                <td>not null</td>
                <td>unique</td>
            </tr>
            <tr>
                <td>Password</td>
                <td>(varchar(100))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Role</td>
                <td>(varchar(20))</td>
                <td>default patient</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Is_active</td>
                <td>(bool)</td>
                <td>default active</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp </td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Deleted at</td>
                <td>(timestamp)</td>
                <td>default null</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Doctors Type table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>id</td>
                <td>int</td>
                <td>not null</td>
                <td>Primary key</td>
            </tr>
            <tr>
                <td>name</td>
                <td>varchar(20)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>created_at</td>
                <td>timestamp</td>
                <td>Current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>updated_at</td>
                <td>timestamp</td>
                <td>Current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>deleted_at</td>
                <td>timestamp</td>
                <td>Default null</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Doctor Details Table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>int</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>User id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Doctor type id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Degree</td>
                <td>varchar(100)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Gender</td>
                <td>varchar(20)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Deleted at</td>
                <td>(timestamp)</td>
                <td>default null  </td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Doctor schedule  table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>int</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>Doctor detail id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Start at</td>
                <td>timestamp</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>End at</td>
                <td>timestamp</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Workdays</td>
                <td>JSON</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Address</td>
                <td>text</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Phone</td>
                <td>varchar(20)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>timestamp</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>timestamp</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>

<h4>Appointment table </h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>User id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Doctor detail id </td>
                <td>(int)</td>
                <td>not null</td>
                <td>foreign key </td>
            </tr>
            <tr>
                <td>Schedule id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Is_appointed</td>
                <td>(bool)</td>
                <td>default false</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Additional Features</h4>
<ul>
<li>GRPC</li>
<li>Postgres</li>
<li>Golang</li>
<li>Microservice Architecture</li>
</ul>
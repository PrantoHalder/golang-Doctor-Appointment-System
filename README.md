##Doctor Appointment System
###Features
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
###Installation
###Requirements
###Tables

#####Users table
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
#####Admin table

#####Doctors Type table
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
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>updated_at</td>
                <td>timestamp</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>deleted_at</td>
                <td>timestamp</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>

#####Doctor Details Table
#####Doctor schedule  table
#####Appointment table


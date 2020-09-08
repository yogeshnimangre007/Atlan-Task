# Atlan-Task

The problem statement for task was as follows 
    We want to offer an implementation through which the user can now stop the long-running task at any given point in time, and can choose to resume or terminate it. This will ensure that the resources like compute/memory/storage/time are used efficiently at our end, and do not go into processing tasks that have already been stopped (and then to roll back the work done post the stop-action)
    
    
   User first give file name to programme which will contain large text data and will start coping it and user will be asked to enter 1 for pausing the process and 2 for terminating the process.Once the process is paused user can resume it by entering 1 and terminate by entering 2 .
    
So what I did is , I assumed a simple process of coping a file content to another and in between this coping  proccess user can pause ,resume and terminate the process. I created a go rouitine  which will do the copying part and two channels that will listen the commands by the user which in turn is in main routine. 
   

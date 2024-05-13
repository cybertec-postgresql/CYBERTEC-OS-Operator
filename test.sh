#!/bin/bash
# Check which Jobs needs to be done

#Define Path
PGBACKREST_PATH=${PGBACKREST_PATH:-'/opt/pgbackrest'}
#source "${PGBACKREST_PATH}/bin/shell_lib.sh"

echo "Start pgBackRest-PreCondition-Check"

if [ "$USE_PGBACKREST" == true ]; then
    echo "Check if RepoHost-Server needs to start"
    if [ "$USE_PGBACKREST" == true ] && [ "$PGBACKREST_SERVER" == true ]; then
        echo "pgBackRest: Start Repo-Host"
        pgbackrest server &
    else
        echo "RepoHost-Server not needed. Skip Step"
    fi

    if [ "$USE_PGBACKREST" == true ] && [ "$PGBACKREST_MODE" == "backup" ]; then
        echo "pgBackRest: Backup-Job found"
        #source "${PGBACKREST_PATH}/bin/backup/start.sh"
        output_success "pgBackRest: Backup-Job completed"
    else
        #For Restore with pgBackrest
        if [ "$RESTORE_ENABLE" == "true" ]; then
            echo "pgBackRest: Restore-Job found"
            #source "${PGBACKREST_PATH}/bin/restore/start.sh"
            output_success "Restore-Job completed"
        else
        echo "Restore not defined - Skip Restore-Step"
            if [ "$RESTORE_BASEBACKUP" == "false" ]; then
                echo "pgBackRest: Backup-Job found"
                export SELECTOR="cluster-name=${SCOPE},spilo-role=master"
                export COMMAND_OPTS="--type=full --stanza=db --repo=1"
                #source "${PGBACKREST_PATH}/bin/backup/start.sh"
                echo "pgBackRest: Update Restore-Configmap"
                configmap="${SCOPE}-pgbackrest-restore"
                kubectl get cm $configmap -o yaml | \
                sed -e 's|restore_basebackup: "false"|restore_basebackup: "true"|' | \
                kubectl apply -f -
                output_success "pgBackRest: Backup-Job completed"
            else
                echo "Basebackup not defined - Skip create basebackup"
            fi
        fi
    fi
else
    echo "pgBackRest not used. Skip Container"
fi
